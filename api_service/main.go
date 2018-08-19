package api_service

import (
	"me/demo/objectStorage/api_service/heartbeat"
	"net/http"
	"me/demo/objectStorage/api_service/objects"
	"me/demo/objectStorage/api_service/locate"
	"log"
	"os"
)

func main() {
	 go heartbeat.ListenHeartbeat()
	 http.HandleFunc("/objects/", objects.Handler)
	 http.HandleFunc("/locate/", locate.Handler)
	 log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
