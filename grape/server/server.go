package server

import (
	"grape/grape/server/api/cluster"
	"grape/grape/server/api/group"
	h "grape/grape/server/api/helpers"
	"grape/grape/server/api/login"
	"grape/grape/server/api/namespace"
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
	cluster.Mount(r)
	login.Mount(r)
	namespace.Mount(r)
	group.Mount(r)
	ui.Mount(r)
	return r
}
