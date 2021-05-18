package service

import (
	"grape/grape/bl"
	h "grape/grape/server/api/helpers"
	"grape/pkg/redislocker"

	"github.com/gin-gonic/gin"
)

func createService(c *gin.Context) {
	s := h.GetSession(c)
	jp := h.GetJsonBody(c)
	name := jp.RequireStr("name")
	code := jp.RequireStr("code")
	port := jp.RequireInt("port")
	protocol := jp.RequireInt("protocol")
	external := jp.OptionalInt("external", 0)
	note := jp.OptionalStr("note", "")
	redislocker.LockP("lock_grape_create_service_"+s.ID, 0, func() {
		service := bl.CreateService(name, code, port, protocol, external, note)
		h.RRJsonObj(c, service)
	})
}
