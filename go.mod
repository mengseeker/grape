module grape

go 1.16

require (
	github.com/Shopify/sarama v1.29.0
	github.com/census-instrumentation/opencensus-proto v0.3.0 // indirect
	github.com/envoyproxy/go-control-plane v0.9.10-0.20210907150352-cf90f659a021
	github.com/envoyproxy/protoc-gen-validate v0.6.1 // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/gofrs/uuid v4.0.0+incompatible
	github.com/golang/protobuf v1.5.2
	github.com/gomodule/redigo v1.8.4
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/influxdata/influxdb1-client v0.0.0-20191209144304-8bf82d3c094d
	github.com/klauspost/compress v1.13.0 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/olivere/elastic/v7 v7.0.26
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.8.1
	go.etcd.io/etcd/api/v3 v3.5.0
	go.etcd.io/etcd/client/v3 v3.5.0
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.19.1
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/net v0.0.0-20210917221730-978cfadd31cf // indirect
	golang.org/x/sys v0.0.0-20210921065528-437939a70204 // indirect
	gomodules.xyz/jsonpatch/v3 v3.0.1
	google.golang.org/genproto v0.0.0-20210921142501-181ce0d877f6
	google.golang.org/grpc v1.41.0
	google.golang.org/protobuf v1.27.1
	k8s.io/api v0.22.2
	k8s.io/apimachinery v0.22.2
	k8s.io/klog/v2 v2.10.0 // indirect
)

replace github.com/envoyproxy/go-control-plane v0.9.8 => github.com/envoyproxy/go-control-plane v0.9.9-0.20201217023817-7fe139bd184a
