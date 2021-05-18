package cluster

import (
	"grape/grape/bl"
	h "grape/grape/server/api/helpers"
	"grape/pkg/redislocker"

	"github.com/gin-gonic/gin"
)

func addlink(c *gin.Context) {
	s := h.GetSession(c)
	jp := h.GetJsonBody(c)
	name := jp.RequireStr("name")
	address := jp.RequireStr("address")
	note := jp.OptionalStr("note", "")
	redislocker.LockP("lock_grape_create_address_"+s.ID, 0, func() {
		r := bl.CreateEtcd(name, address, note)
		h.RRJsonObj(c, r)
	})
}
