package main

import (
	"log"
	"net"
	"fmt"

	"github.com/maulidihsan/flashdeal-webservice/pkg/mongo"
	"github.com/maulidihsan/flashdeal-webservice/pkg/catalog/repository"
	"github.com/maulidihsan/flashdeal-webservice/pkg/catalog/usecase"
	"google.golang.org/grpc"
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

	catalog := repository.NewCatalogCollection(session.Copy(), "crawler", "products")
	us := usecase.NewCatalogUseCase(catalog)
	grpcServer := grpc.NewServer()
	blibli.NewCatalogServerGrpc(grpcServer, us)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}