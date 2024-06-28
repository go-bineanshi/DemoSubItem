package controller

import "github.com/gin-gonic/gin"

type DemoController struct{}

func (dc *DemoController) PingHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Pong",
	})
}
