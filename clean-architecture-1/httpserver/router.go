package httpserver

import (
	"clean-architecture-1/httpserver/controllers"
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/users", controllers.GetUsers)
	router.POST("/users", controllers.CreateUser)

	return router
}
