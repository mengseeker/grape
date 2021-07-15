package service

import (
	"grape/grapeapi/models"
	h "grape/grapeapi/server/api/helpers"

	"github.com/gin-gonic/gin"
)

func serviceInfo(c *gin.Context) {
	id := h.ParamInt(c, "id", 1)
	var service models.Service
	models.ModelService().First(&service, id)
	h.RRJsonObj(c, service)
}
