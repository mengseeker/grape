package login

import (
	"github.com/gin-gonic/gin"
)

func Mount(e *gin.Engine) {
	r := e.Group("/api")
	r.Handle("POST", "/login", normalLogin)
}
