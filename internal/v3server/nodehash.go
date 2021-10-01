package apiv3

import (
	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
)

type nodeHash struct {
}

func (h *nodeHash) ID(node *core.Node) string {
	return node.Id
}
