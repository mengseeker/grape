package service

import "github.com/gin-gonic/gin"

func listService(c *gin.Context) {
	c.String(200, "ok")
}
