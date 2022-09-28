package main

import (
	"assignment-2/config"
	"assignment-2/httpserver"
	"assignment-2/httpserver/controllers"
	"assignment-2/httpserver/repositories/gorm"
	"assignment-2/httpserver/services"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := config.ConnectMysqlGORM()
	if err != nil {
		panic(err)
	}
	orderRepo := gorm.NewOrderRepo(db)
	orderSvc := services.NewOrderSvc(orderRepo)
	orderHandler := controllers.NewOrderController(orderSvc)

	router := gin.Default()

	app := httpserver.NewRouter(router, orderHandler)
	app.Start(":8080")
}
