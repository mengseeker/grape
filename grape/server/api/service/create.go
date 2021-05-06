package service

import (
	"grape/grape/models"
	h "grape/grape/server/api/helpers"
	"grape/pkg/redislocker"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func createService(c *gin.Context) {
	s := h.GetSession(c)
	redislocker.LockP("lock_grape_create_service_"+s.ID, 0, func() {
		name := h.RequirePostFormStr(c, "name")
		code := h.RequirePostFormStr(c, "code")
		port := h.RequirePostFormInt(c, "port")
		protocol := h.RequirePostFormInt(c, "protocol")
		external := h.PostFromInt(c, "external", 0)
		note := h.PostFromStr(c, "note")
		service := models.Service{
			Name:     name,
			Code:     code,
			Port:     port,
			Protocol: protocol,
			External: external,
			Note:     note,
		}
		service.InsertGP(c, boil.Infer())
	})
	h.RRR(c)
}
