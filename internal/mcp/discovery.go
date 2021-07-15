package mcp

import (
	"grape/api"
	"grape/pkg/logger"
	"io"
	"sync"
	"time"

	"google.golang.org/grpc"
)

const (
	updateMaxTk = 10 * time.Second
	updateMinTk = 3 * time.Second
)

type StreamID int

type StreamSet map[StreamID]int

type DiscoveryServer struct {
	api.UnimplementedDiscoveryServiceServer
	Local          *ApiResources
	Streams        map[StreamID]api.DiscoveryService_StreamResourcesServer
	StreamVersions map[StreamID]map[ResourceHash]string
	Watchs         map[ResourceHash]StreamSet
	Components     map[StreamID]*api.Component
	L              logger.Logger

	CurrentStream StreamID
	mux           sync.Mutex
	updateSignal  chan struct{}
	done          chan struct{}
}

func NewDiscoveryServer(l logger.Logger) (*DiscoveryServer, error) {
	s := DiscoveryServer{}
	s.Local = NewApiResources()
	err := s.Local.LoadAll()
	if err != nil {
		return nil, err
	}
	s.Streams = make(map[StreamID]api.DiscoveryService_StreamResourcesServer)
	s.StreamVersions = make(map[StreamID]map[ResourceHash]string)
	s.Watchs = make(map[ResourceHash]StreamSet)
	s.Components = make(map[StreamID]*api.Component)
	s.L = l

	s.updateSignal = make(chan struct{})
	s.done = make(chan struct{})
	return &s, nil
}

func (s *DiscoveryServer) RegisterServer(grpcServer *grpc.Server) {
	api.RegisterDiscoveryServiceServer(grpcServer, s)
}

// TODO 测试防抖
func (s *DiscoveryServer) HandleUpdate() {
	tkMax := time.NewTicker(updateMaxTk)
	tkMin := time.NewTicker(updateMinTk)
	defer tkMax.Stop()
	defer tkMax.Stop()
	for {
		select {
		case <-s.done:
			return
		case <-s.updateSignal:
			// 启动更新
		LOOP_REC:
			for {
				select {
				case <-s.updateSignal:
					tkMin.Reset(updateMinTk)
				case <-tkMax.C:
					break LOOP_REC
				case <-tkMin.C:
					break LOOP_REC
				}
			}
			s.doUpdate()
		}
	}
}

func (s *DiscoveryServer) Cancel() {
	close(s.done)
	close(s.updateSignal)
}

func (s *DiscoveryServer) AddStream(stream api.DiscoveryService_StreamResourcesServer) StreamID {
	s.mux.Lock()
	defer s.mux.Unlock()
	sid := s.CurrentStream
	s.CurrentStream++
	s.Streams[sid] = stream
	s.StreamVersions[sid] = make(map[ResourceHash]string)
	s.L.Infof("stream %d added", sid)
	return sid
}

func (s *DiscoveryServer) RemoveStream(sid StreamID) {
	s.mux.Lock()
	defer s.mux.Unlock()
	delete(s.Streams, sid)
	delete(s.Components, sid)
	delete(s.StreamVersions, sid)
	for _, set := range s.Watchs {
		delete(set, sid)
	}
	s.L.Infof("stream %d removed", sid)
}

func (s *DiscoveryServer) StreamResources(stream api.DiscoveryService_StreamResourcesServer) error {
	sid := s.AddStream(stream)
	defer s.RemoveStream(sid)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			s.L.Warnf("client closed stream %d", sid)
			return err
		} else if err != nil {
			s.L.Errorf("stream %d recv err: %q", sid, err)
			return err
		}
		s.handleDiscovery(sid, req)
	}
}

func (s *DiscoveryServer) handleDiscovery(sid StreamID, req *api.DiscoveryRequest) {
	// register Component and resource watch
	defer s.Update()
	s.mux.Lock()
	defer s.mux.Unlock()
	comm := req.Component
	if s.Components[sid] == nil {
		s.Components[sid] = comm
	}
	resourceHash := Hash(comm.ClusterCode, req.ResourceType)
	if s.Watchs[resourceHash] == nil {
		s.Watchs[resourceHash] = StreamSet{}
	}
	s.Watchs[resourceHash][sid] = 1
	s.StreamVersions[sid][resourceHash] = req.VersionInfo
}

func (s *DiscoveryServer) Update() {
	s.updateSignal <- struct{}{}
}

func (s *DiscoveryServer) doUpdate() {
	defer func() {
		err := recover()
		if err != nil {
			s.L.Errorf("doUpdate panic %v", err)
		}
	}()
	s.mux.Lock()
	defer s.mux.Unlock()
	// 检查资源是否更新
	err := s.Local.CheckUpdate()
	if err != nil {
		s.L.Errorf("update resource fail when checkUpdate: %v", err)
		return
	}
	rhs := s.checkAllResourceVersion()
	for rh, sids := range rhs {
		resp := s.Local.GetResponse(rh)
		for _, sid := range sids {
			err := s.Streams[sid].Send(resp)
			if err != nil {
				s.L.Error("stream %d update resource fail %v", sid, err)
			}
		}
	}

}

// 查找需要更新的stream
func (s *DiscoveryServer) checkAllResourceVersion() map[ResourceHash][]StreamID {
	var updates = map[ResourceHash][]StreamID{}
	for hash, res := range s.Local.Resources {
		sids := []StreamID{}
		version := res.Version
		for sid := range s.Watchs[hash] {
			streamVersion := s.StreamVersions[sid][hash]
			if streamVersion != version {
				sids = append(sids, sid)
			}
		}
		if len(sids) > 0 {
			updates[hash] = sids
		}
	}
	return updates
}
