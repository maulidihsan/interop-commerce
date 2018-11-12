package models
import(
	"time"
)

type Order struct {
	Id string `json:"id"`
	Pembeli User `json:"pembeli"`
	Produk Catalog	`json:"produk"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}