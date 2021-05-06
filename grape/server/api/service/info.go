package service

import (
	"grape/grape/models"
	h "grape/grape/server/api/helpers"

	"github.com/gin-gonic/gin"
)

func serviceInfo(c *gin.Context) {
	id := h.ParamInt(c, "id", 1)
	service, err := models.FindServiceG(c, int64(id))
	if err != nil {
		panic(err)
	}
	h.RRR(c,
		"info", service,
	)
}
