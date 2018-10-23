package main

import (
	"context"
	"log"
	"time"
	"fmt"
	"google.golang.org/grpc"

	"github.com/maulidihsan/flashdeal-webservice/pkg/api"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := api.NewPromoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Call Create
	response, err := c.GetPromo(ctx, &api.Empty{})
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	fmt.Printf("time: %s\n", response.GetTime())
	for _, product := range response.Product{
		fmt.Printf("item: %s, images: %s, stocks: %s, harga: %s\n", product.GetItem(), product.GetImages(), product.GetStocks(), product.GetHarga())
	}
}