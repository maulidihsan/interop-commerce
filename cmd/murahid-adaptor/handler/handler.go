package handler

import (
	"google.golang.org/grpc"
	"github.com/maulidihsan/flashdeal-webservice/pkg/v1"
)

type server struct {}

func NewGrpcServer(gserver *grpc.Server) {

	s := &server{}
	v1.RegisterCatalogServiceServer(gserver, s)
	v1.RegisterOrderServiceServer(gserver, s)
}