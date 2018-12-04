package main

import (
	"github.com/gin-contrib/static"
	"flag"
	"fmt"
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/maulidihsan/interop-commerce/config"
	"github.com/maulidihsan/interop-commerce/pkg/mongo"
	"github.com/maulidihsan/interop-commerce/pkg/order/repository"
	"github.com/maulidihsan/interop-commerce/pkg/order/usecase"
)

type UpdateStatus struct {
	Id string `json:"id"`
	Status string `json:"status"`
}

func main() {
	environment := flag.String("e", "dev", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	c := config.GetConfig()
	session, err := mongo.NewSession(fmt.Sprintf("%s:%s", c.GetString("biru.database.ip"), c.GetString("biru.database.port")), c.GetString("biru.database.dbadmin"), c.GetString("biru.database.user"), c.GetString("biru.database.password"))
	if err != nil {
		log.Printf("%v", err)
		log.Fatalln("unable to connect to mongodb")
	}
	db := repository.NewOrderCollection(session.Copy(), c.GetString("biru.database.name"), "orders")
	usecase := usecase.NewOrderUseCase(db)
	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("./public", true)))
	r.Use(gin.Logger())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/orders", func(c *gin.Context) {
		orders, err := usecase.GetAllOrders()
		if (err != nil) {
			c.AbortWithStatus(500)
			return
		}
		c.JSON(200, gin.H{"orders": orders})
	})
	r.POST("/update-status", func(c *gin.Context) {
		var update UpdateStatus
		c.BindJSON(&update)
		fmt.Println("ID", update.Id)
		res, err := usecase.UpdateStatusOrder(update.Id, update.Status)
		if (err != nil) {
			c.AbortWithStatus(500)
			return
		}
		c.JSON(200, &res)
	})
	r.Run(fmt.Sprintf(":%s", c.GetString("biru.rest.port")))
}
