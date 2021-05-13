package cache

import (
	"sync"
	"time"

	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
)

type nodeInfo struct {
	// node is the constant Envoy node metadata.
	node *core.Node
	seq  int64

	// watches are indexed channels for the response watches and the original requests.
	watches map[int64]ResponseWatch

	// the timestamp of the last watch request
	lastWatchRequestTime time.Time

	// mutex to protect the status fields.
	// should not acquire mutex of the parent cache after acquiring this mutex.
	mu sync.RWMutex
}

// ResponseWatch is a watch record keeping both the request and an open channel for the response.
type ResponseWatch struct {
	// Request is the original request for the watch.
	Request *Request

	// Response is the channel to push responses to.
	Response chan Response
}

// newNodeInfo initializes a status info data structure.
func newNodeInfo(node *core.Node) *nodeInfo {
	out := nodeInfo{
		node:    node,
		watches: make(map[int64]ResponseWatch),
	}
	return &out
}

func (info *nodeInfo) GetNode() *core.Node {
	return info.node
}

func (info *nodeInfo) GetNumWatches() int {
	return len(info.watches)
}

func (info *nodeInfo) GetLastWatchRequestTime() time.Time {
	info.mu.RLock()
	defer info.mu.RUnlock()
	return info.lastWatchRequestTime
}

func (info *nodeInfo) String() string {
	return info.node.Cluster + info.node.Id
}

func (info *nodeInfo) newResponseWatch(req *Request) (int64, chan Response) {
	info.mu.Lock()
	defer info.mu.Unlock()
	info.lastWatchRequestTime = time.Now()
	// allocate capacity 1 to allow one-time non-blocking use
	value := make(chan Response, 1)
	watchID := info.seq
	info.watches[watchID] = ResponseWatch{Request: req, Response: value}
	info.seq++
	return watchID, value
}
