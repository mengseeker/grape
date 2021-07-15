package group

import (
	"grape/grapeapi/bl"
	h "grape/grapeapi/server/api/helpers"

	"github.com/gin-gonic/gin"
)

func createGroup(c *gin.Context) {
	jp := h.GetJsonBody(c)
	name := jp.RequireStr("name")
	code := jp.RequireStr("code")
	namespace := int64(jp.RequireInt("namespace_id"))
	serviceID := int64(jp.RequireInt("service_id"))
	clusterID := int64(jp.RequireInt("cluster_id"))

	version := jp.OptionalStr("version", "")
	lang := jp.OptionalStr("lang", "")
	note := jp.OptionalStr("note", "")

	record := bl.CreateGroup(name, code, version, lang, note, namespace, serviceID, clusterID)
	h.RRJsonObj(c, record)
}
