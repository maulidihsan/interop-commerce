package server

import (
	"fmt"
	"time"
	"log"

	"github.com/maulidihsan/flashdeal-webservice/config"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"github.com/processout/grpc-go-pool"
	"github.com/maulidihsan/flashdeal-webservice/pkg/mongo"
	"github.com/maulidihsan/flashdeal-webservice/cmd/web-service/middlewares"
	// "github.com/maulidihsan/flashdeal-webservice/cmd/web-service/middlewares"
	"github.com/maulidihsan/flashdeal-webservice/cmd/web-service/controllers"
	userRepo "github.com/maulidihsan/flashdeal-webservice/pkg/user/repository"
	userService "github.com/maulidihsan/flashdeal-webservice/pkg/user/usecase"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	c := config.GetConfig()
	session, err := mongo.NewSession(fmt.Sprintf("%s:%s", c.GetString("database.ip"), c.GetString("database.port")), c.GetString("database.dbadmin"), c.GetString("database.user"), c.GetString("database.password"))
	if err != nil {
		log.Printf("%v", err)
		log.Fatalln("unable to connect to mongodb")
	}

	var factory grpcpool.Factory
	factory = func() (*grpc.ClientConn, error) {
		conn, err := grpc.Dial(fmt.Sprintf("%s:%s", c.GetString("murahid.grpc.ip"), c.GetString("murahid.grpc.port")), grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to start gRPC connection: %v", err)
		}
		fmt.Printf("Connected to grpc at %s:%s \n", c.GetString("murahid.grpc.ip"), c.GetString("murahid.grpc.port"))
		return conn, err
	}
	pool, err := grpcpool.New(factory, 5, 5, time.Second)
    if err != nil {
        log.Fatalf("Failed to create gRPC pool: %v", err)
    }

	router.GET("/", new(controllers.HealthController).Status)
	// router.Use(middlewares.AuthMiddleware())

	middlewares.ValidatorInit()
	v1 := router.Group("v1")
	{
		userGroup := v1.Group("users")
		{
			userDB := userRepo.NewUserCollection(session.Copy(), "crawler", "users")
			userUsecase := userService.NewUserUseCase(userDB)
			userController := controllers.NewUserController(userUsecase, middlewares.GetValidator())
			userGroup.POST("/", userController.Register)
			userGroup.POST("/auth", userController.Auth)
			userGroup.Use(middlewares.AuthMiddleware())
			userGroup.GET("/", userController.Get)
		}
		
		catalogGroup := v1.Group("catalogs")
		{
			catalog := controllers.NewCatalogController(pool)
			catalogGroup.Use(middlewares.AuthMiddleware())
			catalogGroup.GET("/", catalog.Get)
			catalogGroup.GET("/:kategori", catalog.BrowseCategory)
		}

		orderGroup := v1.Group("orders")
		{
			order := controllers.NewOrderController(pool, middlewares.GetValidator())
			orderGroup.Use(middlewares.AuthMiddleware())
			orderGroup.GET("/", order.Get)
			orderGroup.POST("/add", order.Add)
		}
	}
	return router
}