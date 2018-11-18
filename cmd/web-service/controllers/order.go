package controllers

import (
	"fmt"
	"github.com/maulidihsan/flashdeal-webservice/pkg/models"
	"encoding/json"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/processout/grpc-go-pool"
	"gopkg.in/go-playground/validator.v9"
	"github.com/maulidihsan/flashdeal-webservice/pkg/v1"
)

type OrderController struct {
	pool *grpcpool.Pool
	validate *validator.Validate
}

func NewOrderController(p *grpcpool.Pool, v *validator.Validate) *OrderController {
	return &OrderController{
		pool: p,
		validate: v,
	}
}

func (p OrderController) Add(c *gin.Context) {
	u := c.MustGet("user").(string)
	var user models.User
	err := json.Unmarshal([]byte(u), &user)
	if(err != nil) {
		c.AbortWithStatus(500)
		return
	}

	var order models.Order
	c.Bind(&order)
	fmt.Println(&order.Produk.Vendor)
	err = p.validate.Struct(&order)
	if(err != nil) {
		var errors []models.ValidationErr
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, models.ValidationErr{
				Tag: err.Tag(),
				Value: err.Value(),
			})
		}
		c.JSON(401, gin.H{"message": "Field validation error", "error": errors})
		c.Abort()
		return
	}
	
	conn, err := p.pool.Get(c)
	defer conn.Close()
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	grpc := v1.NewOrderServiceClient(conn.ClientConn)
	res, err := grpc.CreateOrder(c, &v1.Order{
		Vendor: order.Produk.Vendor,
		Pembeli: &v1.Person{
			Nama: user.Nama,
			Alamat: user.Alamat,
			Email: user.Email,
			Telepon: user.Telepon,
		},
		Barang: &v1.Product{
			Id: order.Produk.Id,
		},
		Kuantitas: order.Kuantitas,
	})
	c.JSON(http.StatusOK, res)
}