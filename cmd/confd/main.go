package main

import (
	"grape/cmd/confd/cmd"
	"grape/pkg/logger"
)

var (
	Version = "0.0.0"
	log     = logger.NewLogger("confd")
)

func main() {
	if err := cmd.NewCmd().Execute(); err != nil {
		log.Fatal(err)
	}
}
