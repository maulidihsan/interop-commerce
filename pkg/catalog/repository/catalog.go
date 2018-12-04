package repository

import models "github.com/maulidihsan/interop-commerce/pkg/models"

type CatalogService interface {
	GetCatalog(keyword string) ([]models.Catalog, error)
	GetByCategory(keyword string) ([]models.Catalog, error)
	GetById(id string) (*models.Catalog, error)
	UpdateCatalog(c []models.Catalog) (*models.Response, error)
}
