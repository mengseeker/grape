package agent

import (
	"context"
	"grape/api/confd"
	"grape/internal/iutils"
	"time"

	"google.golang.org/grpc"
)

var (
	// timeout for first loaddiing of configs
	loadTimeout = 3 * time.Second

	disconveryConn   *grpc.ClientConn
	disconveryClient confd.ConfdServerClient

	// disconveryStreamLock sync.Mutex
	// disconveryStream     confd.ConfdServer_StreamResourcesClient
)

func DiscoveryConfig(ctx context.Context) <-chan *confd.Configs {
	ready := make(chan struct{})
	timeout := time.After(loadTimeout)
	cfChan := make(chan *confd.Configs)
	go func() {
		// dial and get config
		var err error
		dialTimeout, cancel := context.WithTimeout(ctx, time.Second*3)
		defer cancel()
		disconveryConn, err = grpc.DialContext(dialTimeout, config.discoveryAddress, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("unable to connect to discovery server: %v", err)
		}
		go handleDiscovery(ctx, cfChan)
		close(ready)
	}()
	select {
	case <-timeout:
		log.Fatal("timeout to load configs")
	case <-ready:
		log.Infof("discovery server %s connected", config.discoveryAddress)
		break
	}
	return cfChan
}

func handleDiscovery(ctx context.Context, cfs chan<- *confd.Configs) {
	disconveryClient = confd.NewConfdServerClient(disconveryConn)
	for {
		err := discoveryStream(ctx, cfs)
		if err != nil {
			log.Errorf("discoveryStream exit: %v", err)
		} else {
			log.Errorf("discoveryStream exit unexpected")
		}
		time.Sleep(time.Second * 3)
		log.Info("restart discoveryStream")
	}
}

func discoveryStream(ctx context.Context, cfs chan<- *confd.Configs) error {
	streamCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	discovery := &confd.Discovery{
		Service: config.service,
		Node:    iutils.GetNode(),
	}
	disconveryStream, err := disconveryClient.StreamResources(streamCtx, discovery)
	if err != nil {
		return err
	}
	for {
		cf, err := disconveryStream.Recv()
		if err != nil {
			// disconveryStream.CloseSend()
			return err
		}
		// TODO 防抖处理
		cfs <- cf
	}
}
