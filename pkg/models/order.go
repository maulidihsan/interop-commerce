package models

import (
	"time"
)

type Order struct {
	Id string `json:"id"`
	Pembeli *User `json:"pembeli" validate:"omitempty"`
	Produk Catalog	`json:"produk" validate:"required"`
	Kuantitas int32 `json:"pcs" validate:"numeric,required"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}