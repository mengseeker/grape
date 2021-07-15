package h

import (
	"fmt"
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
	} else {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: fmt.Sprintf("未知异常%v", err),
		})
	}
}
