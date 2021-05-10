package service

import (
	"grape/grape/models"
	h "grape/grape/server/api/helpers"

	"github.com/gin-gonic/gin"
)

func serviceInfo(c *gin.Context) {
	id := h.ParamInt(c, "id", 1)
	var service models.Service
	models.ModelService().First(&service, id)
	h.RRJsonObj(c, service)
}
