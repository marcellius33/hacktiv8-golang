package controllers

import (
	"assignment-2/httpserver/controllers/params"
	"assignment-2/httpserver/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

type OrderController struct {
	svc *services.OrderSvc
}

func NewOrderController(svc *services.OrderSvc) *OrderController {
	return &OrderController{
		svc: svc,
	}
}

func (o *OrderController) GetAllOrders(ctx *gin.Context) {
	response := o.svc.GetAllOrders()
	WriteJsonResponse(ctx, response)
}

func (o *OrderController) GetOrder(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	response := o.svc.GetOrder(uint(id))
	WriteJsonResponse(ctx, response)
}

func (o *OrderController) CreateOrder(ctx *gin.Context) {
	var req params.OrderCreateRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = validator.New().Struct(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := o.svc.CreateOrder(&req)
	WriteJsonResponse(ctx, response)
}

func (o *OrderController) UpdateOrder(ctx *gin.Context) {
	var req params.OrderUpdateRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = validator.New().Struct(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := o.svc.UpdateOrder(&req)
	WriteJsonResponse(ctx, response)
}

func (o *OrderController) DeleteOrder(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	response := o.svc.DeleteOrder(uint(id))
	WriteJsonResponse(ctx, response)
}
