package main

import (
	"grape/grape/cmd"
	_ "grape/grape/models"
	"grape/pkg/logger"
)

var log = logger.NewLogger("Grape")

func main() {
	if err := cmd.NewRootCmd().Execute(); err != nil {
		log.Fatal(err)
	}
}
