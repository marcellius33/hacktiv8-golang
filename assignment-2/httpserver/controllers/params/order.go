package params

import "time"

type OrderCreateRequest struct {
	OrderedAt    time.Time                `validate:"required"`
	CustomerName string                   `validate:"required"`
	Items        []OrderItemCreateRequest `validate:"required"`
}

type OrderItemCreateRequest struct {
	ItemCode    string `validate:"required"`
	Description string `validate:"required"`
	Quantity    uint   `validate:"required"`
}

type OrderUpdateRequest struct {
	OrderID      uint                     `validate:"required"`
	OrderedAt    time.Time                `validate:"required"`
	CustomerName string                   `validate:"required"`
	Items        []OrderItemUpdateRequest `validate:"required"`
}

type OrderItemUpdateRequest struct {
	LineItemID  uint   `validate:"required"`
	ItemCode    string `validate:"required"`
	Description string `validate:"required"`
	Quantity    uint   `validate:"required"`
}
