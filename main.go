package main

import (
	"strings"
)

func main() {
	// etcdGetAll()
}

func CutHttpPath(rawPath string) string {
	idx := strings.Index(rawPath, "?")
	if idx >= 0 {
		return rawPath[0:idx]
	}
	return rawPath
}

// func etcdGetAll() {
// 	err := etcdcli.Connect("127.0.0.1:2379")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	kvs, err := etcdcli.PrefixAll(context.Background(), "foo")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println(kvs)
// }
