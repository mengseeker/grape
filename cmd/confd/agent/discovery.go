package agent

import (
	"context"
	"grape/api/v1/confd"
	"time"

	"google.golang.org/grpc"
)

var (
	disconveryConn   *grpc.ClientConn
	disconveryClient confd.ConfdServerClient

	// disconveryStreamLock sync.Mutex
	// disconveryStream     confd.ConfdServer_StreamResourcesClient
)

func DialDiscoveryServer(ctx context.Context) {
	var err error
	dialTimeout, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	disconveryConn, err = grpc.DialContext(dialTimeout, config.discoveryAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("unable to connect to discovery server: %v", err)
	}
	disconveryClient = confd.NewConfdServerClient(disconveryConn)
}

func handleDiscovery(ctx context.Context, cfs chan<- *confd.Configs) {
	for {
		err := discoveryStream(ctx, cfs)
		if err != nil {
			log.Errorf("discoveryStream fail: %v", err)
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
