package namespace

import (
	"grape/grapeapi/bl"
	h "grape/grapeapi/server/api/helpers"

	"github.com/gin-gonic/gin"
)

func create(c *gin.Context) {
	jp := h.GetJsonBody(c)
	record := bl.CreateNs(
		jp.RequireStr("name"),
		jp.RequireStr("code"),
		jp.OptionalStr("note", ""),
		int64(jp.RequireInt("cluster_id")),
	)
	h.RRJsonObj(c, record)
}
