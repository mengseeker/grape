package iutils

import (
	"grape/api/v1/core"
	// "grape/internal/share"
)

func GetNode() *core.Node {
	return &core.Node{
		Ip:      GetLocalIP4(),
		Host:    GetLocalHost(),
	}
}
