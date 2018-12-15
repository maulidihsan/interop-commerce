package handler

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"github.com/maulidihsan/interop-commerce/config"
	"github.com/maulidihsan/interop-commerce/pkg/v1"
)

func (s *server) GetCatalog(ctx context.Context, in *v1.Keyword) (*v1.Products, error) {
	conf := config.GetConfig()
	biruServer, err := grpc.Dial(fmt.Sprintf("%s:%s", conf.GetString("biru.grpc.ip"), conf.GetString("biru.grpc.port")), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer biruServer.Close()

	merahServer, err := grpc.Dial(fmt.Sprintf("%s:%s", conf.GetString("merah.grpc.ip"), conf.GetString("merah.grpc.port")), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer merahServer.Close()

	c1 := v1.NewCatalogServiceClient(biruServer)
	c2 := v1.NewCatalogServiceClient(merahServer)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	response1, err := c1.GetCatalog(ctx, in)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	response2, err := c2.GetCatalog(ctx, in)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}

	response := append(response1.GetProducts(), response2.GetProducts()...)
	res := v1.Products{
		Products: response,
	}
	fmt.Println("Getting Catalog: ", in.GetKeyword())
	return &res, nil
}

func (s *server) GetByCategory(ctx context.Context, in *v1.Keyword) (*v1.Products, error) {
	conf := config.GetConfig()
	biruServer, err := grpc.Dial(fmt.Sprintf("%s:%s", conf.GetString("biru.grpc.ip"), conf.GetString("biru.grpc.port")), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer biruServer.Close()

	merahServer, err := grpc.Dial(fmt.Sprintf("%s:%s", conf.GetString("merah.grpc.ip"), conf.GetString("merah.grpc.port")), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer merahServer.Close()

	c1 := v1.NewCatalogServiceClient(biruServer)
	c2 := v1.NewCatalogServiceClient(merahServer)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	response1, err := c1.GetByCategory(ctx, in)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	response2, err := c2.GetByCategory(ctx, in)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	response := append(response2.GetProducts(), response1.GetProducts()...)
	fmt.Println(response)
	res := v1.Products{
		Products: response,
	}
	fmt.Println("Getting Catalog By Category: ", in.GetKeyword())
	return &res, nil
}

func (s *server) UpdateCatalog(ctx context.Context, in *v1.Products) (*v1.Response, error) {
	conf := config.GetConfig()
	biruServer, err := grpc.Dial(fmt.Sprintf("%s:%s", conf.GetString("biru.grpc.ip"), conf.GetString("biru.grpc.port")), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer biruServer.Close()

	c := v1.NewCatalogServiceClient(biruServer)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	response, err := c.UpdateCatalog(ctx, in)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	fmt.Printf("reply: %v\n", response)
	return response, nil
}

