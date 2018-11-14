package handler

import (
	"google.golang.org/grpc"
	"github.com/maulidihsan/flashdeal-webservice/pkg/v1"
	orderService "github.com/maulidihsan/flashdeal-webservice/pkg/order/usecase"
	catalogService "github.com/maulidihsan/flashdeal-webservice/pkg/catalog/usecase"
	"github.com/maulidihsan/flashdeal-webservice/pkg/models"
)

type server struct {
	catalog catalogService.CatalogUsecase
	order orderService.OrderUsecase
}

func NewOrderServerGrpc(gserver *grpc.Server, orderUsecase orderService.OrderUsecase) {

	orderServer := &server{
		order: orderUsecase,
	}

	v1.RegisterOrderServiceServer(gserver, orderServer)
}

func NewCatalogServerGrpc(gserver *grpc.Server, catalogUsecase catalogService.CatalogUsecase) {

	catalogServer := &server{
		catalog: catalogUsecase,
	}
	v1.RegisterCatalogServiceServer(gserver, catalogServer)
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