package controllers

import (
	"github.com/maulidihsan/flashdeal-webservice/pkg/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/processout/grpc-go-pool"
	"github.com/maulidihsan/flashdeal-webservice/pkg/v1"
)

type CatalogController struct {
	pool *grpcpool.Pool
}

func NewCatalogController(p *grpcpool.Pool) *CatalogController {
	return &CatalogController{p}
}

func (p CatalogController) VendorTagging(in *v1.Product) models.Catalog {
	return models.Catalog{
		Id: in.Id,
		NamaProduk: in.Produk,
		Url: in.Link,
		Gambar: in.Gambar,
		Harga: in.Harga,
		Kategori: in.Kategori,
		Vendor: in.Vendor,
	}
}

func (p CatalogController) Get(c *gin.Context) {
	search := c.DefaultQuery("q", "")
	conn, err := p.pool.Get(c)
	defer conn.Close()
    if err != nil {
        c.JSON(500, gin.H{"message": "Cant Connect to adapter", "error": err})
		c.Abort()
		return
    }
	client := v1.NewCatalogServiceClient(conn.ClientConn)
	data, err := client.GetCatalog(c, &v1.Keyword{
		Keyword: search,
	})
	if(err != nil) {
		c.JSON(404, gin.H{"message": "Not Found", "error": err})
		c.Abort()
		return
	}
	products := make([]models.Catalog, len(data.GetProducts()))
	for i, product := range data.GetProducts(){
		pr := p.VendorTagging(product)
		products[i] = pr
	}
	c.JSON(http.StatusOK, gin.H{"products": products})
}

func (p CatalogController) BrowseCategory(c *gin.Context) {
	category := c.Param("kategori")
	conn, err := p.pool.Get(c)
	defer conn.Close()
    if err != nil {
        c.JSON(500, gin.H{"message": "Cant Connect to adapter", "error": err})
		c.Abort()
		return
    }
	client := v1.NewCatalogServiceClient(conn.ClientConn)
	data, err := client.GetByCategory(c, &v1.Keyword{
		Keyword: category,
	})
	if(err != nil) {
		c.JSON(404, gin.H{"message": "Not Found", "error": err})
		c.Abort()
		return
	}
	products := make([]models.Catalog, len(data.GetProducts()))
	for i, product := range data.GetProducts(){
		pr := p.VendorTagging(product)
		products[i] = pr
	}
	c.JSON(http.StatusOK, gin.H{"products": products})
}