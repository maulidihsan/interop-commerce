package catalog

import (
	"log"
	"fmt"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	models "github.com/maulidihsan/flashdeal-webservice/pkg/models"
	api "github.com/maulidihsan/flashdeal-webservice/pkg/v1"
)

type server struct {
	usecase models.CatalogService
}

func NewCatalogServer(gserver *grpc.Server, service *catalog.CatalogService) {
	catalogServer := &server{
		usecase: service,
	}

	v1.RegisterCatalogServiceServer(gserver, catalogServer)
	reflection.Register(gserver)
}