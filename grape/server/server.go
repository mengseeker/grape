package server

import (
	h "grape/grape/server/api/helpers"
	"grape/grape/server/api/login"
	"grape/grape/server/api/service"
	"grape/grape/server/ui"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.CustomRecovery(h.Recovery))
	r.GET("/health", func(c *gin.Context) {
		c.String(200, "health")
	})
	service.Mount(r)
	login.Mount(r)
	ui.Mount(r)
	return r
}
