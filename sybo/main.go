package main

import (
	"flag"
	"github.com/LinMAD/workshops/sybo/api"
	"github.com/LinMAD/workshops/sybo/storage"
	"log"
	"net/http"
)

const tagMain = "Sybo:"

var apiPort string

func init() {
	flag.StringVar(
		&apiPort,
		"apiPort",
		"8000",
		"HTTP Port for API if game client can be configurable"+
			"Use same port as game client or forward requests via IP tables for debug...",
	)
	flag.Parse()
}

func main() {
	// Create API and provide storage for data saving
	restApi := api.NewAPI(storage.Boot())

	// Listen HTTP Requests and handle requests
	log.Printf("%s Running web server at: http://127.0.0.1:%s/", tagMain, apiPort)
	if serverErr := http.ListenAndServe(":"+apiPort, restApi.Router); serverErr != nil {
		log.Fatal(serverErr.Error())
	}
}
