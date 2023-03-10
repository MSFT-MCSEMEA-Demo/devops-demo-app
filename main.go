package main

import (
	"log"
	"net/http"
	"os"
)

var ver string

func main() {
	rtr := http.DefaultServeMux
	rtr.HandleFunc("/", home{}.handle)
	//not real connectionstring only for demo purposes
	connstring := "DefaultEndpointsProtocol=https;AccountName=ypdemodevstorage;AccountKey=123RRiXSui754qt/QDnTQbZRAFGDH4e1IyNZj4EnKP/R2v5ACoNGuZM9yg1TAmS7lVoc1hWp3luC+AStfNEzZw==;EndpointSuffix=core.windows.net"
	addr := os.Getenv("HTTP_ADDR")

	log.Printf("%s: info: http listen and serve demoapp: %s", ver, addr)
	log.Printf(connstring)
	if err := http.ListenAndServe(addr, rtr); err != nil && err != http.ErrServerClosed {
		log.Printf("%s: error: http listen and serve demoapp: %s", ver, err)
	}
}

type home struct{}

func (h home) handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: info: X-Request-ID: %s\n", ver, r.Header.Get("X-Request-ID"))
	_, _ = w.Write([]byte("Hello world, demoapp image demo tag:" + ver))
}
