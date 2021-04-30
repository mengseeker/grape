package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ok codes
const (
	OkCode = 0
)

// err codes
const (
	AuthFail = iota + 100
)

func R(c *gin.Context, r *Response) {
	c.JSON(http.StatusOK, r)
}

func AuthErr(realm string) *Response {
	if realm == "" {
		realm = "Authorization Required"
	}
	return &Response{Code: AuthFail, Message: realm, Data: nil}
}
