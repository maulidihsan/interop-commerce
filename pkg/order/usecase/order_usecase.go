package usecase

import (
	models "github.com/maulidihsan/flashdeal-webservice/pkg/models"
	repository "github.com/maulidihsan/flashdeal-webservice/pkg/order/repository"
)

type OrderUsecase interface {
	GetOrders(id string) ([]models.Order, error)
	UpdateStatusOrder(id string, status string) (*models.Response, error)
	CreateOrder(*models.Order) (*models.Response, error)
	DeleteOrder(id string) (*models.Response, error)
	GetAllOrders() ([]models.Order, error)
}

type orderUsecase struct {
	orderRepos *repository.OrderCollection
}

func(o *orderUsecase) GetOrders(id string) ([]models.Order, error) {
	return o.orderRepos.GetOrders(id)
}

func(o *orderUsecase) UpdateStatusOrder(id string, status string) (*models.Response, error) {
	return o.orderRepos.UpdateStatusOrder(id, status)
}

func(o *orderUsecase) CreateOrder(newOrder *models.Order) (*models.Response, error) {
	return o.orderRepos.CreateOrder(newOrder)
}

func(o *orderUsecase) DeleteOrder(id string) (*models.Response,error) {
	return o.orderRepos.DeleteOrder(id)
}

func(o *orderUsecase) GetAllOrders() ([]models.Order, error) {
	return o.orderRepos.GetAllOrders()
}

func NewOrderUseCase(o *repository.OrderCollection) OrderUsecase {
	return &orderUsecase{o}
}