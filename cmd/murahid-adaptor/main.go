package main

import (
	"net"
	"log"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	murahid "github.com/maulidihsan/flashdeal-webservice/cmd/murahid-adaptor/handler"
)
func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", 8888))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	murahid.NewCatalogServerGrpc(grpcServer)
	murahid.NewOrderServerGrpc(grpcServer)
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}