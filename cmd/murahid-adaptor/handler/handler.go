package handler

import (
	"google.golang.org/grpc"
	"github.com/maulidihsan/flashdeal-webservice/pkg/v1"
	orderService "github.com/maulidihsan/flashdeal-webservice/pkg/order/usecase"
	catalogService "github.com/maulidihsan/flashdeal-webservice/pkg/catalog/usecase"
)

type server struct {
	catalog catalogService.CatalogUsecase
	order orderService.OrderUsecase
}

func NewOrderServerGrpc(gserver *grpc.Server) {

	orderServer := &server{}
	v1.RegisterOrderServiceServer(gserver, orderServer)
}

func NewCatalogServerGrpc(gserver *grpc.Server) {

	catalogServer := &server{}
	v1.RegisterCatalogServiceServer(gserver, catalogServer)
}