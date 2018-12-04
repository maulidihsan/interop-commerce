package models

import (
	"time"
	"github.com/maulidihsan/interop-commerce/pkg/v1"
)

type Catalog struct {
	Id string `json:"id" validate:"required"`
	Vendor v1.Vendor `json:"vendor"`
	NamaProduk string `json:"product_name"`
	Url string	`json:"link"`
	Gambar string	`json:"image"`
	Harga int32 `json:"harga"`
	Kategori string	`json:"kategori"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}
