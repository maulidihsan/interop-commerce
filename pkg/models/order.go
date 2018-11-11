package models

type Order struct {
	Id string `json:"id"`
	Pembeli User `json:"pembeli"`
	Produk Catalog	`json:"produk"`
}

type OrderService interface {
	AddOrder(c *Order) error
	GetById(id string) (*Order,error)
}