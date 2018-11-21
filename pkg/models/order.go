package models

import (
	"time"
	"github.com/maulidihsan/flashdeal-webservice/pkg/v1"
)

type Order struct {
	Id string `json:"id"`
	Vendor v1.Vendor `json:"vendor"`
	Pembeli *User `json:"pembeli" validate:"omitempty"`
	Produk Catalog	`json:"produk" validate:"required"`
	Kuantitas int32 `json:"pcs" validate:"numeric,required"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}