package main

import (
	"grape/logtrans/cmd"
	"grape/pkg/logger"
)

var log = logger.NewLogger("logtrans")

func main() {
	err := cmd.NewRootCmd().Execute()
	if err != nil {
		log.Fatal(err)
	}
}
