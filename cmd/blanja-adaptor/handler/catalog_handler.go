package handler

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"github.com/maulidihsan/flashdeal-webservice/pkg/v1"
)


func (s *server) GetCatalog(ctx context.Context, in *v1.Keyword) (*v1.Products, error) {
	blibliServer, err := grpc.Dial("localhost:7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer blibliServer.Close()

	c := v1.NewCatalogServiceClient(blibliServer)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	response, err := c.GetCatalog(ctx, in)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	fmt.Println("Getting Catalog: ", in.GetKeyword())
	return response, nil
}

func (s *server) GetByCategory(ctx context.Context, in *v1.Keyword) (*v1.Products, error) {
	blibliServer, err := grpc.Dial("localhost:7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer blibliServer.Close()

	c := v1.NewCatalogServiceClient(blibliServer)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	response, err := c.GetByCategory(ctx, in)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	fmt.Println("Getting Catalog By Category: ", in.GetKeyword())
	return response, nil
}

func (s *server) UpdateCatalog(ctx context.Context, in *v1.Products) (*v1.Response, error) {
	blibliServer, err := grpc.Dial("localhost:7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer blibliServer.Close()

	c := v1.NewCatalogServiceClient(blibliServer)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	response, err := c.UpdateCatalog(ctx, in)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	fmt.Printf("reply: %v\n", response)
	return response, nil
}

