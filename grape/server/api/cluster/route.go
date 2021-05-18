package cluster

import (
	h "grape/grape/server/api/helpers"

	"github.com/gin-gonic/gin"
)

func Mount(e *gin.Engine) {
	r := e.Group("/api/cluster")
	r.Use(h.Auth)
	r.POST("/create_link", addlink)
	r.POST("/create", create)
	r.GET("/list", List)
}
