package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-bineanshi/DemoSubItem/routes"
)

func main() {
	r := gin.Default()
	routes.Route(r)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
