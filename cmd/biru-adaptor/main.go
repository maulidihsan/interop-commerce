package main

import (
	"log"
	"net"
	"flag"
	"os"
	"fmt"

	"github.com/maulidihsan/interop-commerce/pkg/mongo"
	catalogRepo "github.com/maulidihsan/interop-commerce/pkg/catalog/repository"
	catalogService "github.com/maulidihsan/interop-commerce/pkg/catalog/usecase"
	orderRepo "github.com/maulidihsan/interop-commerce/pkg/order/repository"
	orderService "github.com/maulidihsan/interop-commerce/pkg/order/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	biru "github.com/maulidihsan/interop-commerce/cmd/biru-adaptor/handler"
	"github.com/maulidihsan/interop-commerce/config"
)


func main() {
	environment := flag.String("e", "dev", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	c := config.GetConfig()


	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", c.GetInt32("biru.grpc.port")))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	session, err := mongo.NewSession(fmt.Sprintf("%s:%s", c.GetString("biru.database.ip"), c.GetString("biru.database.port")), c.GetString("biru.database.dbadmin"), c.GetString("biru.database.user"), c.GetString("biru.database.password"))
	if err != nil {
		log.Printf("%v", err)
		log.Fatalln("unable to connect to mongodb")
	}

	catalog := catalogRepo.NewCatalogCollection(session.Copy(), c.GetString("biru.database.name"), "catalogs")
	catalogUseCase := catalogService.NewCatalogUseCase(catalog)

	order := orderRepo.NewOrderCollection(session.Copy(), c.GetString("biru.database.name"), "orders")
	orderUseCase := orderService.NewOrderUseCase(order)
	grpcServer := grpc.NewServer()
	biru.NewServerGrpc(grpcServer, catalogUseCase, orderUseCase)
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
