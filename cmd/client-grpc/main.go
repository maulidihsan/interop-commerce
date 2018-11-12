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
	conn, err := grpc.Dial("localhost:7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := api.NewOrderServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Call Create
	response, err := c.DeleteOrder(ctx, &api.OrderId{
		OrderId: "5be9cdb4305904c27447ef81",
	})
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	fmt.Printf("reply: %s\n", response.GetMessage())
}