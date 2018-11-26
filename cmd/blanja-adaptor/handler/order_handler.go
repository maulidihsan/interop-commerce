package handler

import (
	"fmt"
	"time"
	"log"
	"context"

	"google.golang.org/grpc"
	"github.com/maulidihsan/flashdeal-webservice/pkg/v1"
)

func (s *server) GetOrders(ctx context.Context, in *v1.UserId) (*v1.Orders, error) {
	blibliServer, err := grpc.Dial("localhost:7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer blibliServer.Close()

	c := v1.NewOrderServiceClient(blibliServer)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := c.GetOrders(ctx, in)
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}
	fmt.Println("Getting user orders")
	return response, nil
}

func (s *server) CreateOrder(ctx context.Context, in *v1.Order) (*v1.Response, error) {
	blibliServer, err := grpc.Dial("localhost:7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer blibliServer.Close()

	c := v1.NewOrderServiceClient(blibliServer)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	fmt.Println("Creating Order for ", in.Vendor)
	if (in.Vendor.String() == "BLIBLI") {
		response, err := c.CreateOrder(ctx, in)
		if err != nil {
			log.Fatalf("Request failed: %v", err)
		}
		return response, nil
	} else {
		res := &v1.Response{
			Success: false,
			Code: 400,
			Message: "Wrong Service Destination",
		}
		return res, nil
	}
}

func (s *server) UpdateStatusOrder(ctx context.Context, in *v1.Update) (*v1.Response, error) {
	blibliServer, err := grpc.Dial("localhost:7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer blibliServer.Close()

	c := v1.NewOrderServiceClient(blibliServer)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if (in.Vendor.String() == "BLIBLI") {
		response, err := c.UpdateStatusOrder(ctx, in)
		if err != nil {
			log.Fatalf("Request failed: %v", err)
		}
		return response, nil
	} else {
		res := &v1.Response{
			Success: false,
			Code: 400,
			Message: "Wrong Service Destination",
		}
		return res, nil
	}
}

func (s *server) DeleteOrder(ctx context.Context, in *v1.OrderId) (*v1.Response, error) {
	blibliServer, err := grpc.Dial("localhost:7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer blibliServer.Close()

	c := v1.NewOrderServiceClient(blibliServer)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if (in.Vendor.String() == "BLIBLI") {
		response, err := c.DeleteOrder(ctx, in)
		if err != nil {
			log.Fatalf("Request failed: %v", err)
		}
		return response, nil
	} else {
		res := &v1.Response{
			Success: false,
			Code: 400,
			Message: "Wrong Service Destination",
		}
		return res, nil
	}
}