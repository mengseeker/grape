package main

import (
	"grape/cmd/confd/agent"
	"grape/pkg/logger"
)

var (
	Version = "0.0.0"
	log     = logger.NewLogger("confd")
)

func main() {
	if err := agent.NewAgentCmd().Execute(); err != nil {
		log.Fatal(err)
	}
}
