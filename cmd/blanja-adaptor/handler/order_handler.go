package handler

import (
	"fmt"
	"time"
	"log"
	"context"

	"google.golang.org/grpc"
	"github.com/maulidihsan/interop-commerce/config"
	"github.com/maulidihsan/interop-commerce/pkg/v1"
)

func (s *server) GetOrders(ctx context.Context, in *v1.UserId) (*v1.Orders, error) {
	conf := config.GetConfig()
	biruServer, err := grpc.Dial(fmt.Sprintf("%s:%s", conf.GetString("biru.grpc.ip"), conf.GetString("biru.grpc.port")), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer biruServer.Close()

	c := v1.NewOrderServiceClient(biruServer)
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
	defer biruServer.Close()

	c1 := v1.NewOrderServiceClient(biruServer)
	c2 := v1.NewOrderServiceClient(merahServer)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	fmt.Println("Creating Order for ", in.Vendor)
	if (in.Vendor.String() == "BLIBLI") {
		response, err := c1.CreateOrder(ctx, in)
		if err != nil {
			log.Fatalf("Request failed: %v", err)
		}
		return response, nil
	} else if (in.Vendor.String() == "BUKALAPAK") {
		response, err := c2.CreateOrder(ctx, in)
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
	defer biruServer.Close()

	c1 := v1.NewOrderServiceClient(biruServer)
	c2 := v1.NewOrderServiceClient(merahServer)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if (in.Vendor.String() == "BLIBLI") {
		response, err := c1.UpdateStatusOrder(ctx, in)
		if err != nil {
			log.Fatalf("Request failed: %v", err)
		}
		return response, nil
	} else if (in.Vendor.String() == "BUKALAPAK") {
		response, err := c2.UpdateStatusOrder(ctx, in)
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
	defer biruServer.Close()

	c1 := v1.NewOrderServiceClient(biruServer)
	c2 := v1.NewOrderServiceClient(merahServer)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if (in.Vendor.String() == "BLIBLI") {
		response, err := c1.DeleteOrder(ctx, in)
		if err != nil {
			log.Fatalf("Request failed: %v", err)
		}
		return response, nil
	} else if (in.Vendor.String() == "BUKALAPAK") {
		response, err := c2.DeleteOrder(ctx, in)
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
