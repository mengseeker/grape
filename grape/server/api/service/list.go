package service

import (
	"grape/grape/models"
	h "grape/grape/server/api/helpers"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func listService(c *gin.Context) {
	page := h.ParamInt(c, "page", 1)
	list, err := models.Services(
		qm.Offset(page*h.PerPage-h.PerPage),
		qm.Limit(h.PerPage),
		qm.OrderBy("id desc"),
	).AllG(c)
	if err != nil {
		panic(err)
	}
	searchInfo := map[string]interface{}{
		"page":     page,
		"per_page": h.PerPage,
		"sort":     "id desc",
	}
	h.RRR(c,
		"search_info", searchInfo,
		"list", list,
	)
}
