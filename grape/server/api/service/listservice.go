package service

import (
	"grape/grape/models"
	h "grape/grape/server/api/helpers"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func listService(c *gin.Context) {
	page := h.ParamInt(c, "page")
	if page <= 0 {
		page = 1
	}
	list, err := models.Services(
		qm.Offset(page*h.PerPage-h.PerPage),
		qm.Limit(h.PerPage),
	).AllG(c)
	if err != nil {
		panic(err)
	}
	searchInfo := map[string]interface{}{}
	h.RRR(c,
		"search_info", searchInfo,
		"list", list,
	)
}
