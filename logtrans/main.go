package main

import "grape/logtrans/server"

func main() {
	go server.ServeAccessLog(":9412")
	go server.ServeZikpin(":9411")
	<-make(chan int)
}
