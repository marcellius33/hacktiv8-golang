package httpserver

import (
	"assignment-2/httpserver/controllers"
	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
	order  *controllers.OrderController
}

func NewRouter(router *gin.Engine, order *controllers.OrderController) *Router {
	return &Router{
		router: router,
		order:  order,
	}
}

func (r *Router) Start(port string) {
	r.router.GET("/ping", controllers.HealthCheck)

	r.router.GET("/orders", r.order.GetAllOrders)
	r.router.GET("/orders/:id", r.order.GetOrder)
	r.router.POST("/orders", r.order.CreateOrder)
	r.router.PUT("/orders/:id", r.order.UpdateOrder)
	r.router.DELETE("/orders/:id", r.order.DeleteOrder)

	r.router.Run(port)
}
