package h

import (
	"fmt"
	"grape/pkg/session"

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

func GetSession(c *gin.Context) *session.Session {
	val, exists := c.Get("session")
	if !exists {
		panic(fmt.Errorf("session not found, use auth first"))
	}
	return val.(*session.Session)
}
