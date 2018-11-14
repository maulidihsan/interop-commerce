package repository

import models "github.com/maulidihsan/flashdeal-webservice/pkg/models"

type OrderService interface {
	GetOrders(email string) ([]models.Order, error)
	CreateOrder(order *models.Order) (*models.Response, error)
	UpdateStatusOrder(id string, status string) (*models.Response, error)
	DeleteOrder(id string) (*models.Response, error)
}