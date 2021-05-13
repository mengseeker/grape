package server

import (
	"grape/logtrans/server/forward"
	"grape/logtrans/server/logs"
	"io/ioutil"
	"log"
	"net/http"
)

func ServeZikpin(address string) {
	http.HandleFunc("/api/v1/spans", zipkinV1Handle)
	http.HandleFunc("/api/v2/spans", zipkinV2Handle)
	log.Println("ServeZikpin at: ", address)
	log.Fatal(http.ListenAndServe(address, nil))
}

func zipkinV1Handle(rw http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s\n", r.Method, r.URL.Path)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("read body err: %v", err)
	}
	entity := logs.Trace{Log: string(body)}
	forward.AddEnvoyTraceLog(&entity)
}

func zipkinV2Handle(rw http.ResponseWriter, r *http.Request) {
	// todo
}
