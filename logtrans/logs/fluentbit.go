package logs

import (
	"encoding/json"
	"grape/pkg/util"
)

type FluentBitLog struct {
	Timestamp float64 `json:"@timestamp"`
	Path      string  `json:"_path"`
	Log       string  `json:"log"`
}

func GetFluentBitMessageLog(m *Message) []byte {
	l := FluentBitLog{}
	json.Unmarshal(m.Val, &l)
	return util.Str2bytes(l.Log)
}
