package main

import (
	"log"
	"net/http"

	"ggentile.com/websockets1/internal/handlers"
)

func main() {
	mux := routes()

	log.Println("Starting channnel Listener")
	go handlers.ListenToWsChannel()

	log.Println("Starting web server on port 8080")

	_ = http.ListenAndServe(":8080", mux)
}
