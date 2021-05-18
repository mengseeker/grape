package cluster

import (
	"grape/grape/bl"
	h "grape/grape/server/api/helpers"
	"grape/pkg/redislocker"

	"github.com/gin-gonic/gin"
)

func create(c *gin.Context) {
	s := h.GetSession(c)
	jp := h.GetJsonBody(c)
	name := jp.RequireStr("name")
	code := jp.RequireStr("code")
	deployType := jp.RequireStr("deploy_type")
	etcdLinkID := jp.RequireInt("etcd_link")
	note := jp.OptionalStr("note", "")

	redislocker.LockP("lock_grape_create_cluster_"+s.ID, 0, func() {
		r := bl.CreateCluster(name, code, note, deployType, etcdLinkID)
		h.RRJsonObj(c, r)
	})
}
