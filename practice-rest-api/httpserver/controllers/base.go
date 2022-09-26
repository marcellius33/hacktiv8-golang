package controllers

import (
	"github.com/gin-gonic/gin"
	"practice-rest-api/httpserver/controllers/views"
)

func WriteJsonResponse(ctx *gin.Context, resp *views.Response) {
	ctx.JSON(resp.Status, resp)
}
