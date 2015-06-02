package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	log.SetFlags(log.Lshortfile)

	proxyPort := os.Getenv("TURBO_PORT")
	if proxyPort == "" {
		proxyPort = "8228"
	}

	handlerPort := os.Getenv("HANDLER_PORT")
	if handlerPort == "" {
		handlerPort = "8227"
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8003"
	}

	go startProxy(proxyPort)
	go startOrderHandler(handlerPort)

	staticServer := http.FileServer(http.Dir("static"))
	log.Println("Static server listening on port " + httpPort)
	log.Fatal(http.ListenAndServe(":"+httpPort, staticServer))
}
