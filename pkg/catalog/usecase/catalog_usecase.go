package usecase

import (
	models "github.com/maulidihsan/flashdeal-webservice/pkg/models"
	repository "github.com/maulidihsan/flashdeal-webservice/pkg/catalog/repository"
)

type CatalogUsecase interface {
	GetCatalog(keyword string) ([]models.Catalog, error)
	UpdateCatalog(products []models.Catalog) (*models.Response, error)
}

type catalogUsecase struct {
	catalogRepos *repository.CatalogCollection
}

func (c *catalogUsecase) GetCatalog(keyword string) ([]models.Catalog, error) {
	return c.catalogRepos.GetCatalog(keyword)
}

func (c *catalogUsecase) UpdateCatalog(catalogs []models.Catalog) (*models.Response, error) {
	return c.catalogRepos.UpdateCatalog(catalogs)
}

func NewCatalogUseCase(c *repository.CatalogCollection) CatalogUsecase {
	return &catalogUsecase{c}
}