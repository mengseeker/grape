package main

import (
	"grape/extauth/cmd"
	"grape/pkg/logger"
)

var log = logger.NewLogger("Auth")

func main() {
	if err := cmd.NewRootCmd().Execute(); err != nil {
		log.Fatal(err)
	}
}
