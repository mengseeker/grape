package main

import (
	"context"
	"grape/api/v1/confd"
	"grape/internal/iutils"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.DialContext(context.Background(), ":15010", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	service := "dev/demo"
	version := iutils.NewVersion()
	cli := confd.NewApiServerClient(conn)
	cf := confd.Configs{
		Version:     version,
		Service:     service,
		FileConfigs: []*confd.FileConfig{{Path: "/tmp/grape_demo.yaml", Content: "aa: 1"}},
	}
	_, err = cli.Set(context.Background(), &confd.SetRequest{
		ServerConfig: &confd.ServerConfig{
			Service: service,
			Default: &cf,
		}})
	if err != nil {
		log.Fatal(err)
	}
}
