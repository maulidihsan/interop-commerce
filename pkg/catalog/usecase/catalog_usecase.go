package usecase

import (
	models "github.com/maulidihsan/flashdeal-webservice/pkg/models"
	repository "github.com/maulidihsan/flashdeal-webservice/pkg/catalog/repository"
)

type CatalogUsecase interface {
	GetCatalog(keyword string) ([]models.Catalog, error)
	GetByCategory(keyword string) ([]models.Catalog, error)
	GetById(id string) (*models.Catalog, error)
	UpdateCatalog(products []models.Catalog) (*models.Response, error)
}

type catalogUsecase struct {
	catalogRepos *repository.CatalogCollection
}

func (c *catalogUsecase) GetCatalog(keyword string) ([]models.Catalog, error) {
	return c.catalogRepos.GetCatalog(keyword)
}

func (c *catalogUsecase) GetByCategory(keyword string) ([]models.Catalog, error) {
	return c.catalogRepos.GetByCategory(keyword)
}

func (c *catalogUsecase) GetById(id string) (*models.Catalog, error) {
	return c.catalogRepos.GetById(id)
}

func (c *catalogUsecase) UpdateCatalog(catalogs []models.Catalog) (*models.Response, error) {
	return c.catalogRepos.UpdateCatalog(catalogs)
}

func NewCatalogUseCase(c *repository.CatalogCollection) CatalogUsecase {
	return &catalogUsecase{c}
}