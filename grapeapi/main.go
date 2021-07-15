package main

import (
	"grape/grapeapi/cmd"
	_ "grape/grapeapi/models"
	"grape/pkg/logger"
)

var log = logger.NewLogger("Grape")

func main() {
	if err := cmd.NewRootCmd().Execute(); err != nil {
		log.Fatal(err)
	}
}
