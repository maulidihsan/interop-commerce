package main

import (
	"log"
	"net"
	"fmt"

	"github.com/maulidihsan/flashdeal-webservice/pkg/mongo"
	catalogRepo "github.com/maulidihsan/flashdeal-webservice/pkg/catalog/repository"
	catalogService "github.com/maulidihsan/flashdeal-webservice/pkg/catalog/usecase"
	orderRepo "github.com/maulidihsan/flashdeal-webservice/pkg/order/repository"
	orderService "github.com/maulidihsan/flashdeal-webservice/pkg/order/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	blibli "github.com/maulidihsan/flashdeal-webservice/cmd/blibli-adaptor/handler"
)


func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	session, err := mongo.NewSession("127.0.0.1:27017")
	if err != nil {
		log.Printf("%v", err)
		log.Fatalln("unable to connect to mongodb")
	}

	catalog := catalogRepo.NewCatalogCollection(session.Copy(), "crawler", "products")
	catalogUseCase := catalogService.NewCatalogUseCase(catalog)

	order := orderRepo.NewOrderCollection(session.Copy(), "crawler", "orders")
	orderUseCase := orderService.NewOrderUseCase(order)
	grpcServer := grpc.NewServer()
	blibli.NewCatalogServerGrpc(grpcServer, catalogUseCase)
	blibli.NewOrderServerGrpc(grpcServer, orderUseCase)
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}