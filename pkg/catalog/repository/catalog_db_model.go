package repository

import (
	"time"
	//"fmt"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"github.com/maulidihsan/interop-commerce/pkg/models"
)

type catalogModel struct {
	Id bson.ObjectId `bson:"_id,omitempty"`
	NamaProduk string `bson:"product_name"`
	Url string	`bson:"link"`
	Gambar string	`bson:"image"`
	Harga int32 `bson:"harga"`
	Kategori string	`bson:"kategori"`
	CreatedAt *time.Time `bson:"created_at"`
}

func catalogModelIndex() mgo.Index {
	return mgo.Index{
		Key: []string{"url"},
		Unique: true,
		DropDups: true,
		Background: true,
		Sparse: true,
	}
}

func newCatalogModel(c *models.Catalog) catalogModel {
	return catalogModel{
		NamaProduk: c.NamaProduk,
		Url: c.Url,
		Gambar: c.Gambar,
		Harga: c.Harga,
		Kategori: c.Kategori,
		CreatedAt: c.CreatedAt,
	}
}

type CatalogArray []catalogModel

func(c CatalogArray) toCatalogs() []models.Catalog {
	var catalogs []models.Catalog
	for _, product := range c {
		//fmt.Println(product)
		catalogs = append(catalogs, models.Catalog{
			Id: product.Id.Hex(),
			NamaProduk: product.NamaProduk,
			Url: product.Url,
			Gambar: product.Gambar,
			Harga: product.Harga,
			Kategori: product.Kategori,
			CreatedAt: product.CreatedAt,
		})
	}
	return catalogs
}

func(u *catalogModel) toCatalog() *models.Catalog {
	return &models.Catalog{
		Id: u.Id.Hex(),
		NamaProduk: u.NamaProduk,
		Url: u.Url,
		Gambar: u.Gambar,
		Harga: u.Harga,
		Kategori: u.Kategori,
		CreatedAt: u.CreatedAt,
	}
}
