package main

import (
	"me/demo/objectStorage/data_service/heartbeat"
	"me/demo/objectStorage/data_service/locate"
	"net/http"
	"log"
	"os"
	"me/demo/objectStorage/data_service/objects"
)

func main() {
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}