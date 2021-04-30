package helpers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParamInt(c *gin.Context, key string) int {
	r, _ := strconv.Atoi(c.Param("page"))
	return r
}
