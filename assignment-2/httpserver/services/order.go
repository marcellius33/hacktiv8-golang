package services

import (
	"assignment-2/httpserver/controllers/params"
	"assignment-2/httpserver/controllers/views"
	"assignment-2/httpserver/repositories"
	"assignment-2/httpserver/repositories/models"
	"database/sql"
	"strings"
)

type OrderSvc struct {
	repo repositories.OrderRepo
}

func NewOrderSvc(repo repositories.OrderRepo) *OrderSvc {
	return &OrderSvc{
		repo: repo,
	}
}

func (o *OrderSvc) GetAllOrders() *views.Response {
	orders, err := o.repo.GetOrders()
	if err != nil {
		if err == sql.ErrNoRows {
			return views.DataNotFound(err)
		}
		return views.InternalServerError(err)
	}

	return views.SuccessFindAllResponse(parseModelToOrderGetAll(orders), "GET_ALL_ORDERS")
}

func (o *OrderSvc) GetOrder(id uint) *views.Response {
	order, err := o.repo.GetOrder(id)
	if err != nil {
		return views.InternalServerError(err)
	}
	return views.SuccessFindResponse(parseModelToOrder(order), "GET_ORDER")
}

func (o *OrderSvc) CreateOrder(req *params.OrderCreateRequest) *views.Response {
	order := parseCreateRequestToModel(req)

	err := o.repo.CreateOrder(order)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return views.DataConflict(err)
		}
		return views.InternalServerError(err)
	}
	return views.SuccessCreateResponse(order, "CREATE_ORDER")
}

func (o *OrderSvc) UpdateOrder(req *params.OrderUpdateRequest) *views.Response {
	newOrder := parseUpdateRequestToModel(req)

	order, err := o.repo.GetOrder(newOrder.ID)
	if err != nil {
		return views.InternalServerError(err)
	}
	order.CustomerName = newOrder.CustomerName
	order.OrderedAt = newOrder.OrderedAt
	order.Items = newOrder.Items

	newOrder, err = o.repo.UpdateOrder(order)
	if err != nil {
		return views.InternalServerError(err)
	}

	return views.SuccessUpdateResponse(newOrder, "UPDATE_ORDER")
}

func (o *OrderSvc) DeleteOrder(id uint) *views.Response {
	order, err := o.repo.GetOrder(id)
	if err != nil {
		return views.InternalServerError(err)
	}
	err = o.repo.DeleteOrder(order)
	if err != nil {
		return views.InternalServerError(err)
	}

	return views.SuccessDeleteResponse(order, "DELETE_ORDER")
}

func parseModelToOrderGetAll(mod *[]models.Order) *[]views.OrderGetAll {
	var o []views.OrderGetAll
	for _, order := range *mod {
		var i []views.OrderItemGetAll
		for _, item := range order.Items {
			i = append(i, views.OrderItemGetAll{
				LineItemID:  item.ID,
				ItemCode:    item.Code,
				Description: item.Description,
				Quantity:    item.Quantity,
			})
		}
		o = append(o, views.OrderGetAll{
			ID:           order.ID,
			CustomerName: order.CustomerName,
			OrderedAt:    order.OrderedAt,
			Items:        i,
		})
	}
	return &o
}

func parseModelToOrder(mod *models.Order) *views.OrderGetAll {
	var i []views.OrderItemGetAll
	for _, item := range mod.Items {
		i = append(i, views.OrderItemGetAll{
			LineItemID:  item.ID,
			ItemCode:    item.Code,
			Description: item.Description,
			Quantity:    item.Quantity,
		})
	}
	o := views.OrderGetAll{
		ID:           mod.ID,
		CustomerName: mod.CustomerName,
		OrderedAt:    mod.OrderedAt,
		Items:        i,
	}
	return &o
}

func parseCreateRequestToModel(req *params.OrderCreateRequest) *models.Order {
	var i []models.Item
	for _, item := range req.Items {
		i = append(i, models.Item{
			Code:        item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
		})
	}
	return &models.Order{
		CustomerName: req.CustomerName,
		OrderedAt:    req.OrderedAt,
		Items:        i,
	}
}

func parseUpdateRequestToModel(req *params.OrderUpdateRequest) *models.Order {
	var i []models.Item
	for _, item := range req.Items {
		i = append(i, models.Item{
			ID:          item.LineItemID,
			Code:        item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
			OrderID:     req.OrderID,
		})
	}
	return &models.Order{
		ID:           req.OrderID,
		CustomerName: req.CustomerName,
		OrderedAt:    req.OrderedAt,
		Items:        i,
	}
}
