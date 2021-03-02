package main

import (
	"grape"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.FS(grape.StaticFS)))
	http.ListenAndServe(":3000", nil)
}
