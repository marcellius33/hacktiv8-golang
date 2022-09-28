package models

type Item struct {
	ID          uint   `json:"lineItemId" gorm:"primary_key;autoIncrement"`
	Code        string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint
}
