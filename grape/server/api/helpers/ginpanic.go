package helpers

import "github.com/gin-gonic/gin"

func Recovery(c *gin.Context, err interface{}) {
	e, ok := err.(error)
	if ok {
		RR(c, 1, e.Error())
	}
}
