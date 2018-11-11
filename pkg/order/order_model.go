package order

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"github.com/maulidihsan/flashdeal-webservice/pkg/catalog"
)

type orderModel struct {
	Id bson.ObjectId `bson:"_id,omitempty"`
	NamaProduk string `bson:"nama_produk"`
	Url string	`bson:"url"`
	Gambar string	`bson:"gambar"`
	Harga int32 `bson:"harga"`
	Kategori string	`bson:"kategori"`
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

func newCatalogModel(c *Catalog) *catalogModel {
	return &catalogModel{
		NamaProduk: c.NamaProduk,
		Url: c.Url,
		Gambar: c.Gambar,
		Harga: c.Harga,
		Kategori: c.Kategori,
	}
}

func(c *catalogModel) toCatalog() *Catalog {
	return &Catalog{
		Id: c.Id.Hex(),
		Url: c.Url,
		Gambar: c.Gambar,
		Harga: c.Harga,
		Kategori: c.Kategori,
	}
}