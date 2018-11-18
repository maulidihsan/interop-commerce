package main

import (
	"log"
	"net"
	"flag"
	"os"
	"fmt"

	"github.com/maulidihsan/flashdeal-webservice/pkg/mongo"
	catalogRepo "github.com/maulidihsan/flashdeal-webservice/pkg/catalog/repository"
	catalogService "github.com/maulidihsan/flashdeal-webservice/pkg/catalog/usecase"
	orderRepo "github.com/maulidihsan/flashdeal-webservice/pkg/order/repository"
	orderService "github.com/maulidihsan/flashdeal-webservice/pkg/order/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	blibli "github.com/maulidihsan/flashdeal-webservice/cmd/blibli-adaptor/handler"
	"github.com/maulidihsan/flashdeal-webservice/config"
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


	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", c.GetInt32("blibli.port")))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	session, err := mongo.NewSession(fmt.Sprintf("%s:%s", c.GetString("database.ip"), c.GetString("database.port")), c.GetString("database.dbadmin"), c.GetString("database.user"), c.GetString("database.password"))
	if err != nil {
		log.Printf("%v", err)
		log.Fatalln("unable to connect to mongodb")
	}

	catalog := catalogRepo.NewCatalogCollection(session.Copy(), "crawler", "products")
	catalogUseCase := catalogService.NewCatalogUseCase(catalog)

	order := orderRepo.NewOrderCollection(session.Copy(), "crawler", "orders")
	orderUseCase := orderService.NewOrderUseCase(order)
	grpcServer := grpc.NewServer()
	blibli.NewServerGrpc(grpcServer, catalogUseCase, orderUseCase)
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}