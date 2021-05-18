package h

import (
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

type JsonBody struct {
	// raw string
	jr gjson.Result
}

func GetJsonBody(c *gin.Context) *JsonBody {
	bs, _ := io.ReadAll(c.Request.Body)
	j := JsonBody{jr: gjson.ParseBytes(bs)}
	return &j
}

func (j *JsonBody) OptionalStr(key string, defalutVal string) string {
	r := j.jr.Get(key)
	if !r.Exists() {
		return defalutVal
	}
	return r.String()
}

func (j *JsonBody) RequireStr(key string) string {
	r := j.jr.Get(key)
	if !r.Exists() {
		panic(fmt.Errorf("param %s is required", key))
	}
	return r.String()
}

func (j *JsonBody) OptionalInt(key string, defalutVal int) int {
	r := j.jr.Get(key)
	if !r.Exists() {
		return defalutVal
	}
	return int(r.Int())
}

func (j *JsonBody) RequireInt(key string) int {
	r := j.jr.Get(key)
	if !r.Exists() {
		panic(fmt.Errorf("param %s is required", key))
	}
	return int(r.Int())
}
