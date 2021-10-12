package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
	handleSign()
}

func handleSign() {
	c := make(chan os.Signal, 3)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for s := range c {
		println("sign", s, "ignored")
	}
}
