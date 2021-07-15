package cluster

import (
	"grape/grapeapi/bl"
	h "grape/grapeapi/server/api/helpers"
	"grape/pkg/redislocker"

	"github.com/gin-gonic/gin"
)

func create(c *gin.Context) {
	s := h.GetSession(c)
	jp := h.GetJsonBody(c)
	name := jp.RequireStr("name")
	cluster_type := jp.RequireStr("cluster_type")
	note := jp.OptionalStr("note", "")

	redislocker.LockP("lock_grape_create_cluster_"+s.ID, 0, func() {
		r := bl.CreateCluster(name, cluster_type, note)
		h.RRJsonObj(c, r)
	})
}
