package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"practice-rest-api/httpserver/controllers/params"
	"practice-rest-api/httpserver/services"
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
