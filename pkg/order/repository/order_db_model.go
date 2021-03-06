package repository

import (
	"time"
	"gopkg.in/mgo.v2/bson"
	"github.com/maulidihsan/interop-commerce/pkg/models"
)

type Pembeli struct {
	// Id bson.ObjectId `bson:"id_user"`
	Nama string `bson:"nama"`
	Alamat string `bson:"alamat"`
	Telepon string `bson:"telepon"`
	Email string `bson:"email"`
}

type Product struct {
	Id bson.ObjectId `bson:"id_product"`
	NamaProduk string `bson:"product_name"`
	Url string	`bson:"link"`
	Gambar string	`bson:"image"`
	Harga int32 `bson:"harga"`
	Kategori string	`bson:"kategori"`
}

type orderModel struct {
	Id bson.ObjectId `bson:"_id,omitempty"`
	Pembeli Pembeli `bson:"pembeli"`
	Produk Product	`bson:"produk"`
	Kuantitas int32 `bson:"pcs"`
	Status string `bson:"status"`
	CreatedAt time.Time `bson:"created_at"`
}

func newOrderModel(o *models.Order) (*orderModel) {
	
	return &orderModel{
		Pembeli: Pembeli{
			Nama: o.Pembeli.Nama,
			Alamat: o.Pembeli.Alamat,
			Telepon: o.Pembeli.Telepon,
			Email: o.Pembeli.Email,
		},
		Produk: Product{
			Id: bson.ObjectIdHex(o.Produk.Id),
			NamaProduk: o.Produk.NamaProduk,
			Url: o.Produk.Url,
			Gambar: o.Produk.Gambar,
			Harga: o.Produk.Harga,
			Kategori: o.Produk.Kategori,
		},
		Kuantitas: o.Kuantitas,
		Status: o.Status,
		CreatedAt: time.Now(),
	}
}

type OrderArray []orderModel

func(o OrderArray) toOrders() []models.Order {
	var orders []models.Order
	for _, order := range o {
		orders = append(orders, models.Order{
			Id: order.Id.Hex(),
			Pembeli: &models.User{
				Nama: order.Pembeli.Nama,
				Alamat: order.Pembeli.Alamat,
				Telepon: order.Pembeli.Telepon,
				Email: order.Pembeli.Email,
			},
			Produk: models.Catalog{
				Id: order.Produk.Id.Hex(),
				NamaProduk: order.Produk.NamaProduk,
				Url: order.Produk.Url,
				Gambar: order.Produk.Gambar,
				Harga: order.Produk.Harga,
				Kategori: order.Produk.Kategori,
			},
			Kuantitas: order.Kuantitas,
			Status: order.Status,
			CreatedAt: order.CreatedAt,
		})
	}
	return orders
}
