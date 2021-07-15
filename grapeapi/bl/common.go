package bl

import "grape/grapeapi/models"

type SearchInfo struct {
	Page    int    `json:"page"`
	PerPage int    `json:"per_page"`
	Total   int64  `json:"total"`
	Order   string `json:"order"`
}

var db = models.GetDB

type gromParam = map[string]interface{}

func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}
