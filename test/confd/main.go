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
	version := iutils.NewVersion()
	cli := confd.NewApiServerClient(conn)
	cf := confd.Configs{
		Version: version,
		FileConfigs: []*confd.FileConfig{
			{Path: "/tmp/grape_demo.yaml", Content: "aa: 1"},
			{Path: "logs/not_exist.yaml", Content: "aa: 1"},
		},
		RestartType: confd.Configs_Kill,
		// RestartType:    confd.Configs_WriteFiles,
		// RestartType:    confd.Configs_Command,
		RunCmd:         "top",
		RestartCommand: "apk add htop",
	}
	_, err = cli.Set(context.Background(), &confd.SetRequest{
		ServerConfig: &confd.ServerConfig{
			Namespace: "demo",
			Service:   "testapp",
			Default:   &cf,
		}})
	if err != nil {
		log.Fatal(err)
	}
}
