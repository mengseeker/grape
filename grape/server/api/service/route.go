package service

import (
	h "grape/grape/server/api/helpers"

	"github.com/gin-gonic/gin"
)

func Mount(e *gin.Engine) {
	r := e.Group("/api/service")
	r.Use(h.Auth)
	r.GET("/list", listService)
	r.GET("/info", serviceInfo)
	r.POST("/create", createService)
}
