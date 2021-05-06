package helpers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParamInt(c *gin.Context, key string, defaultVal int) int {
	r, err := strconv.Atoi(c.Param(key))
	if err != nil {
		return defaultVal
	}
	return r
}

func RequireParamInt(c *gin.Context, key string) int {
	val := c.Param("page")
	if val == "" {
		panic(fmt.Errorf("param %s is required", key))
	}
	r, err := strconv.Atoi(val)
	if err != nil {
		panic(fmt.Errorf("param %s err: %v", key, err))
	}
	return r
}

func PostFromStr(c *gin.Context, key string) string {
	return c.PostForm(key)
}

func RequirePostFormStr(c *gin.Context, key string) string {
	val := c.PostForm(key)
	if val == "" {
		panic(fmt.Errorf("param %s is required", key))
	}
	return val
}

func RequirePostFormInt(c *gin.Context, key string) int {
	val := c.PostForm(key)
	if val == "" {
		panic(fmt.Errorf("param %s is required", key))
	}
	r, err := strconv.Atoi(val)
	if err != nil {
		panic(fmt.Errorf("param %s err: %v", key, err))
	}
	return r
}


func PostFromInt(c *gin.Context, key string, defaultVal int) int {
	r, err := strconv.Atoi(c.PostForm(key))
	if err != nil {
		return defaultVal
	}
	return r
}

