package login

import (
	h "grape/grape/server/api/helpers"
	"grape/pkg/session"

	"github.com/gin-gonic/gin"
)

func normalLogin(c *gin.Context) {
	jp := h.GetJsonBody(c)
	username := jp.RequireStr("username")
	password := jp.RequireStr("password")
	if username == "admin" && password == "admin123" {
		s := session.NewSession(1)
		s.Save()
		h.RRR(c,
			"token", s.ID,
		)
	}
}
