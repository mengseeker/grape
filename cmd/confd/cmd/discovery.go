package cmd

import (
	"context"
	"fmt"
	confdv1 "grape/api/v1/confd"
	"time"

	"google.golang.org/grpc"
)

var (
	disconveryConn   *grpc.ClientConn
	disconveryClient confdv1.ConfdServerClient

	// disconveryStreamLock sync.Mutex
	// disconveryStream     confd.ConfdServer_StreamResourcesClient
)

func runDiscovery(discoveryChan chan<- *confdv1.Configs) {
	if err := dialDiscoveryServer(context.Background()); err != nil {
		log.Fatal(err)
	}
	handleDiscovery(context.Background(), discoveryChan)
}


func dialDiscoveryServer(ctx context.Context) error {
	var err error
	dialTimeout, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	disconveryConn, err = grpc.DialContext(dialTimeout, config.discoveryAddress, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("unable to connect to discovery server: %v", err)
	}
	disconveryClient = confdv1.NewConfdServerClient(disconveryConn)
	return nil
}

func handleDiscovery(ctx context.Context, cfs chan<- *confdv1.Configs) {
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

func discoveryStream(ctx context.Context, cfs chan<- *confdv1.Configs) error {
	streamCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	discovery := &confdv1.Discovery{
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
