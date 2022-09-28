package mysql

import (
	"assignment-2/httpserver/repositories"
	"assignment-2/httpserver/repositories/models"
	"database/sql"
)

type orderRepo struct {
	db *sql.DB
}

func NewOrderRepo(db *sql.DB) repositories.OrderRepo {
	return &orderRepo{
		db: db,
	}
}

func (o *orderRepo) GetOrders() (*[]models.Order, error) {

}

func (o *orderRepo) GetOrder(id uint) (*models.Order, error) {

}

func (o *orderRepo) CreateOrder(order *models.Order) error {

}

func (o *orderRepo) UpdateOrder(order *models.Order) (*models.Order, error) {

}

func (o *orderRepo) DeleteOrder(order *models.Order) error {

}
