package cluster

import (
	"grape/grape/bl"
	h "grape/grape/server/api/helpers"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	list, searchInfo := bl.ListClusters()
	h.RRR(c,
		"search_info", searchInfo,
		"list", list,
	)
}
