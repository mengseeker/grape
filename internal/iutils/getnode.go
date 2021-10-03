package iutils

import (
	"grape/api/core"
	"grape/internal/share"
)

func GetNode() *core.Node {
	return &core.Node{
		Service: share.GetService(),
		Ip:      GetLocalIP4(),
		Host:    GetLocalHost(),
	}
}
