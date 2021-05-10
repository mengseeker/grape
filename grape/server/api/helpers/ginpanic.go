package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Recovery(c *gin.Context, err interface{}) {
	e, ok := err.(error)
	if ok {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: e.Error(),
		})
	}
}
