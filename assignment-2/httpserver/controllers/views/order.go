package views

import "time"

type OrderGetAll struct {
	ID           uint              `json:"orderId"`
	CustomerName string            `json:"customerName"`
	OrderedAt    time.Time         `json:"orderedAt"`
	Items        []OrderItemGetAll `json:"items"`
}

type OrderItemGetAll struct {
	LineItemID  uint   `json:"lineItemId"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
}
