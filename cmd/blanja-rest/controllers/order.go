package controllers

import (
	"fmt"
	"time"
	"github.com/maulidihsan/interop-commerce/pkg/models"
	"encoding/json"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/processout/grpc-go-pool"
	"gopkg.in/go-playground/validator.v9"
	"github.com/maulidihsan/interop-commerce/pkg/v1"
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

func vendorTagging(in *v1.Order) models.Order {
	return models.Order{
		Vendor: in.Vendor,
		Id: in.Id,
		Pembeli: &models.User{
			Nama: in.Pembeli.Nama,
			Alamat: in.Pembeli.Alamat,
			Telepon: in.Pembeli.Telepon,
			Email: in.Pembeli.Email,
		},
		Produk: models.Catalog{
			Id: in.Barang.Id,
			NamaProduk: in.Barang.Produk,
			Url: in.Barang.Link,
			Gambar: in.Barang.Gambar,
			Harga: in.Barang.Harga,
			Kategori: in.Barang.Kategori,
		},
		Kuantitas: in.Kuantitas,
		Status: in.Status,
		CreatedAt: time.Unix(in.CreatedAt, 0),
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

func(p OrderController) Get(c *gin.Context) {
	u := c.MustGet("user").(string)
	var user models.User
	err := json.Unmarshal([]byte(u), &user)
	if(err != nil) {
		c.AbortWithStatus(500)
		return
	}
	conn, err := p.pool.Get(c)
	defer conn.Close()
	if err != nil {
		c.JSON(500, gin.H{"message": "Can't connect to adapter", "error": err})
		c.Abort()
		return
	}
	client := v1.NewOrderServiceClient(conn.ClientConn)
	data, err := client.GetOrders(c, &v1.UserId{
		Id: user.Email,
	})
	if err != nil {
		c.JSON(404, gin.H{"message": "Not Found", "error": err})
		c.Abort()
		return
	}
	orders := make([]models.Order, len(data.GetOrders()))
	for i, order := range data.GetOrders(){
		orders[i] = vendorTagging(order)
	}
	c.JSON(http.StatusOK, gin.H{"orders": orders})
}
