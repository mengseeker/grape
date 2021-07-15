package h

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
	OkCode         = 0
	SuccessMessage = "success"
)

// err codes
const (
	AuthFail = iota + 100
)

func R(c *gin.Context, r *Response) {
	c.JSON(http.StatusOK, r)
}

func RR(c *gin.Context, code int, message string, stringAndValData ...interface{}) {
	data := map[string]interface{}{}
	for i := 0; i < len(stringAndValData)-1; i += 2 {
		data[stringAndValData[i].(string)] = stringAndValData[i+1]
	}
	r := Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
	R(c, &r)
}

func RRJsonObj(c *gin.Context, obj interface{}) {
	r := Response{
		Code:    OkCode,
		Message: SuccessMessage,
		Data:    obj,
	}
	R(c, &r)
}

func RRR(c *gin.Context, stringAndValData ...interface{}) {
	RR(c, OkCode, SuccessMessage, stringAndValData...)
}

func AuthErr(realm string) *Response {
	if realm == "" {
		realm = "Authorization Required"
	}
	return &Response{Code: AuthFail, Message: realm, Data: nil}
}
