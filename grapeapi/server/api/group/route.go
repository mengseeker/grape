package group

import (
	h "grape/grapeapi/server/api/helpers"

	"github.com/gin-gonic/gin"
)

func Mount(e *gin.Engine) {
	r := e.Group("/api/group")
	r.Use(h.Auth)
	r.POST("/create", createGroup)
}
