package controllers

import (
	"clean-architecture-1/httpserver/controllers/params"
	"clean-architecture-1/httpserver/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func CreateUser(ctx *gin.Context) {
	var req params.UserCreateRequest

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

	response := services.CreateUser(&req)

	WriteJsonResponse(ctx, response)
}

func GetUsers(ctx *gin.Context) {
	response := services.GetUsers()

	WriteJsonResponse(ctx, response)
}
