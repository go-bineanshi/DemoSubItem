package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/go-bineanshi/DemoSubItem/controllers"
)

func Route(engine gin.IRouter) {
	demoController := new(controller.DemoController)
	engine.GET("/ping", demoController.PingHandler)
}
