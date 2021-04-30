package server

import (
	"grape/grape/server/api/service"
	"grape/grape/server/ui"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.String(200, "health")
	})
	service.Mount(r)
	ui.Mount(r)
	return r
}
