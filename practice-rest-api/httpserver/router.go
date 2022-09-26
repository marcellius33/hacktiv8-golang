package httpserver

import (
	"github.com/gin-gonic/gin"
	"practice-rest-api/httpserver/controllers"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/user", controllers.CreateUser)

	return router
}
