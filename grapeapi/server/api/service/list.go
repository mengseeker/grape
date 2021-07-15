package service

import (
	"grape/grapeapi/bl"
	h "grape/grapeapi/server/api/helpers"
	"strconv"

	"github.com/gin-gonic/gin"
)

func listService(c *gin.Context) {
	page := h.ParamInt(c, "page", 1)
	limit := h.PerPage
	offset := page*limit - limit
	params := map[string]interface{}{}
	h.ParamExist(c, "id", func(id string) {
		params["id"], _ = strconv.Atoi(id)
	})
	list, searchInfo := bl.SearchService(limit, offset, params)
	h.RRR(c,
		"search_info", searchInfo,
		"list", list,
	)
}
