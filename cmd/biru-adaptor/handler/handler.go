package handler

import (
	"google.golang.org/grpc"
	"github.com/maulidihsan/interop-commerce/pkg/v1"
	orderService "github.com/maulidihsan/interop-commerce/pkg/order/usecase"
	catalogService "github.com/maulidihsan/interop-commerce/pkg/catalog/usecase"
	"github.com/maulidihsan/interop-commerce/pkg/models"
)

type server struct {
	catalog catalogService.CatalogUsecase
	order orderService.OrderUsecase
}


func NewServerGrpc(gserver *grpc.Server, catalogUsecase catalogService.CatalogUsecase, orderUsecase orderService.OrderUsecase) {

	s := &server{
		catalog: catalogUsecase,
		order: orderUsecase,
	}
	v1.RegisterCatalogServiceServer(gserver, s)
	v1.RegisterOrderServiceServer(gserver, s)
}


func (s *server) transformResponseRPC(ar *models.Response) *v1.Response {
	if ar == nil {
		return nil
	}
	res := &v1.Response{
		Success: ar.Success,
		Code: ar.Code,
		Message: ar.Message,
	}
	return res
}
