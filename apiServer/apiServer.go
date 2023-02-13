package main

import (
	"github.janlei/objio/apiServer/heartbeat"
	"github.janlei/objio/apiServer/locate"
	"github.janlei/objio/apiServer/objects"
	"github.janlei/objio/apiServer/temp"
	"github.janlei/objio/apiServer/versions"
	"log"
	"net/http"
	"os"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	http.HandleFunc("/versions/", versions.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
