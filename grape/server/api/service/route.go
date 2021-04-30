package service

import (
	"grape/grape/server/api/helpers"

	"github.com/gin-gonic/gin"
)

func Mount(e *gin.Engine) {
	r := e.Group("/api/service")
	r.Use(helpers.Auth)
	r.GET("/list", listService)
}
