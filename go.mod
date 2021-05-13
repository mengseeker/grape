module grape

go 1.16

require (
	github.com/cncf/udpa/go v0.0.0-20210322005330-6414d713912e // indirect
	github.com/envoyproxy/go-control-plane v0.9.8
	github.com/envoyproxy/protoc-gen-validate v0.6.1 // indirect
	github.com/gin-gonic/gin v1.7.1
	github.com/gofrs/uuid v3.2.0+incompatible
	github.com/gomodule/redigo v1.8.4
	github.com/lib/pq v1.10.1 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.6.3
	github.com/tidwall/gjson v1.7.5
	go.etcd.io/etcd/api/v3 v3.5.0-alpha.0
	go.etcd.io/etcd/client/v3 v3.5.0-alpha.0
	go.uber.org/zap v1.16.0
	golang.org/x/crypto v0.0.0-20210506145944-38f3c27a63bf // indirect
	golang.org/x/lint v0.0.0-20201208152925-83fdc39ff7b5 // indirect
	golang.org/x/net v0.0.0-20210508051633-16afe75a6701 // indirect
	golang.org/x/sys v0.0.0-20210507161434-a76c4d0a0096 // indirect
	golang.org/x/tools v0.1.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013
	google.golang.org/grpc v1.35.0-dev
	google.golang.org/protobuf v1.26.0
	gorm.io/driver/postgres v1.1.0
	gorm.io/gorm v1.21.9
)

replace github.com/envoyproxy/go-control-plane v0.9.8 => github.com/envoyproxy/go-control-plane v0.9.9-0.20201217023817-7fe139bd184a
