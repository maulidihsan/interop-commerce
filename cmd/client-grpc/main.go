package main

import (
	"context"
	"log"
	"time"
	"fmt"
	"google.golang.org/grpc"

	api "github.com/maulidihsan/flashdeal-webservice/pkg/v1"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := api.NewOrderServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Call Create
	response, err := c.CreateOrder(ctx, &api.Order{
		Vendor: 0,
		Pembeli: &api.Person{
			Nama: "Maulid Ihsan",
			Alamat: "JAKAL",
			Email: "maulid@maul.id",
			Telepon: "010101",
		},
		Barang: &api.Product{
			Produk: "Asus Zenfone",
			Link: "https://blibli.com/",
			Gambar: "https://blibli.com/",
			Harga: 7000000,
			Kategori: "handphone",
		},
	})
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	fmt.Printf("reply: %s\n", response.GetMessage())
}