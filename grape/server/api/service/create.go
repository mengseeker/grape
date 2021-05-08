package service

import (
	"grape/grape/bl"
	h "grape/grape/server/api/helpers"
	"grape/pkg/redislocker"

	"github.com/gin-gonic/gin"
)

func createService(c *gin.Context) {
	s := h.GetSession(c)
	name := h.RequirePostFormStr(c, "name")
	code := h.RequirePostFormStr(c, "code")
	port := h.RequirePostFormInt(c, "port")
	protocol := h.RequirePostFormInt(c, "protocol")
	external := h.PostFromInt(c, "external", 0)
	note := h.PostFromStr(c, "note")
	redislocker.LockP("lock_grape_create_service_"+s.ID, 0, func() {
		service := bl.CreateService(name, code, port, protocol, external, note)
		h.RRJsonObj(c, service)
	})
}
