package repositories

import "assignment-2/httpserver/repositories/models"

type OrderRepo interface {
	GetOrders() (*[]models.Order, error)
	GetOrder(id uint) (*models.Order, error)
	CreateOrder(order *models.Order) error
	UpdateOrder(order *models.Order) (*models.Order, error)
	DeleteOrder(order *models.Order) error
}
