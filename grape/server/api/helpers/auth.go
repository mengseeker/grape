package helpers

import (
	"grape/grape/pkg/session"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	token := c.Request.Header.Get("Token")
	if token == "" {
		R(c, AuthErr(""))
		c.Abort()
		return
	}
	s, err := session.Find(token)
	if err != nil {
		R(c, AuthErr(err.Error()))
		c.Abort()
		return
	}
	c.Set("session", s)
}
