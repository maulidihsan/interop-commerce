package models

import "time"

type Catalog struct {
	Id string `json:"id"`
	NamaProduk string `json:"product_name"`
	Url string	`json:"link"`
	Gambar string	`json:"image"`
	Harga int32 `json:"harga"`
	Kategori string	`json:"kategori"`
	CreatedAt time.Time `json:"created_at"`
}