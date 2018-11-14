package repository

import models "github.com/maulidihsan/flashdeal-webservice/pkg/models"

type CatalogService interface {
	GetCatalog(keyword string) ([]models.Catalog, error)
	UpdateCatalog(c []models.Catalog) (*models.Response, error)
}