package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type DemoController struct{}

func (dc *DemoController) PingHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": fmt.Sprintf("Pong at %s", time.Now().String()),
	})
}
