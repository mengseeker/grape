package cache

import (
	"context"
	"grape/pkg/logger"
	"sync"

	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	// discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	cachev3 "github.com/envoyproxy/go-control-plane/pkg/cache/v3"
)

type Request = cachev3.Request
type Response = cachev3.Response
type RawResponse = cachev3.RawResponse
// type Cache = cachev3.Cache
type Node = core.Node

type v3Cache struct {
	log   logger.Logger
	mu    sync.RWMutex
	nodes map[string]*nodeInfo
}

func New(l logger.Logger) *v3Cache {
	c := v3Cache{
		log: l,
	}
	return &c
}

func (cache *v3Cache) GetNodeInfo(node *Node) *nodeInfo {
	nodeID := node.Id
	info, ok := cache.nodes[nodeID]
	if !ok {
		info = newNodeInfo(node)
		cache.nodes[nodeID] = info
	}
	return info
}

func (c *v3Cache) Fetch(context.Context, *Request) (Response, error)

// CreateWatch returns a watch for an xDS request.
func (cache *v3Cache) CreateWatch(request *Request) (chan Response, func()) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	info := cache.GetNodeInfo(request.Node)
	watchID, value := info.newResponseWatch(request)

	if cache.log != nil {
		cache.log.Infof("open watch %d for %s%v from nodeID %q, version %q", watchID,
			request.TypeUrl, request.ResourceNames, info, request.VersionInfo)
	}
	return value, cache.cancelWatch(info, watchID)
}

func (c *v3Cache) GetVersion(resourceType string, info *nodeInfo) string

func (c *v3Cache) cancelWatch(info *nodeInfo, watchID int64) func()
