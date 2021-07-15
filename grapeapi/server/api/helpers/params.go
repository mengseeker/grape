package h

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParamExist(c *gin.Context, key string, caller ...func(val string)) bool {
	val, ok := c.GetQuery(key)
	if ok {
		for _, cal := range caller {
			cal(val)
		}
	}
	return ok
}

func ParamInt(c *gin.Context, key string, defaultVal int) int {
	r, err := strconv.Atoi(c.Query(key))
	if err != nil {
		return defaultVal
	}
	return r
}

func RequireParamInt(c *gin.Context, key string) int {
	val := c.Query("page")
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
