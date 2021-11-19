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
			{Path: "logs/not_exist2.yaml", Content: "aa: 1"},
		},
		RestartType: confd.RestartType_Kill,
		// RestartType:    confd.Configs_WriteFiles,
		// RestartType:    confd.Configs_Command,
		RunCmd:         "sh -c tail -f",
		RestartCommand: "apk add htop",
	}
	_, err = cli.Set(context.Background(), &confd.ApiRequest{
		ProjectName: "demo/order",
		Project: &confd.Project{
			Name: "demo/order",
			GroupConfigs: map[string]*confd.Configs{
				"default": &cf,
			},
		}})
	if err != nil {
		log.Fatal(err)
	}
}
