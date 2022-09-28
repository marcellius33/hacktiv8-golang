package gorm

import (
	"assignment-2/httpserver/repositories"
	"assignment-2/httpserver/repositories/models"
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
)

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) repositories.OrderRepo {
	return &orderRepo{
		db: db,
	}
}

func (o *orderRepo) GetOrders() (*[]models.Order, error) {
	var orders []models.Order
	err := o.db.Model(&models.Order{}).Preload("Items").Find(&orders).Error
	if err != nil {
		return nil, err
	}

	if len(orders) == 0 {
		return nil, sql.ErrNoRows
	}

	return &orders, nil
}

func (o *orderRepo) GetOrder(id uint) (*models.Order, error) {
	var order models.Order
	err := o.db.Model(&models.Order{}).Preload("Items").First(&order, id).Error
	if err != nil {
		return nil, err
	}
	fmt.Println(order.Items)
	return &order, nil
}

func (o *orderRepo) CreateOrder(order *models.Order) error {
	return o.db.Create(order).Error
}

func (o *orderRepo) UpdateOrder(order *models.Order) (*models.Order, error) {
	err := o.db.Save(&order).Error
	return order, err
}

func (o *orderRepo) DeleteOrder(order *models.Order) error {
	o.db.Delete(&order.Items)
	return o.db.Delete(&order).Error
}
