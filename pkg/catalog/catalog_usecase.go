package catalog

import (
	"github.com/maulidihsan/flashdeal-webservice/pkg/models"
)

type CatalogUseCase interface {
	GetCatalog(filters []string) ([]*models.Catalog, error)
	UpdateCatalog(product []*models.Catalog) (*models.Response, error)
}