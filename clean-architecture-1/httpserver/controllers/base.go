package controllers

import (
	"clean-architecture-1/httpserver/controllers/views"
	"github.com/gin-gonic/gin"
)

func WriteJsonResponse(ctx *gin.Context, resp *views.Response) {
	ctx.JSON(resp.Status, resp)
}
