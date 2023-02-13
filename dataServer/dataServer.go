package main

import (
	"github.janlei/objio/apiServer/temp"
	"github.janlei/objio/dataServer/heartbeat"
	"github.janlei/objio/dataServer/locate"
	"github.janlei/objio/dataServer/objects"
	"log"
	"net/http"
	"os"
)

func main() {
	locate.CollectObjects()
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
