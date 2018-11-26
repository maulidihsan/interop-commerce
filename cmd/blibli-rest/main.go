package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/maulidihsan/flashdeal-webservice/config"
	"github.com/maulidihsan/flashdeal-webservice/pkg/mongo"
	"github.com/maulidihsan/flashdeal-webservice/pkg/order/repository"
	"github.com/maulidihsan/flashdeal-webservice/pkg/order/usecase"
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
	session, err := mongo.NewSession(fmt.Sprintf("%s:%s", c.GetString("blibli.database.ip"), c.GetString("blibli.database.port")), c.GetString("blibli.database.dbadmin"), c.GetString("blibli.database.user"), c.GetString("blibli.database.password"))
	if err != nil {
		log.Printf("%v", err)
		log.Fatalln("unable to connect to mongodb")
	}

	db := repository.NewOrderCollection(session.Copy(), c.GetString("blibli.database.name"), "orders")
	usecase := usecase.NewOrderUseCase(db)
	r := gin.Default()
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
	r.Run(fmt.Sprintf(":%s", c.GetString("blibli.rest.port")))
}
