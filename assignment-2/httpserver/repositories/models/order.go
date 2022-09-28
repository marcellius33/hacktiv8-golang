package models

import (
	"time"
)

type Order struct {
	ID           uint      `json:"orderId" gorm:"primary_key;autoIncrement"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `json:"items" gorm:"foreignKey:ID"`
}
