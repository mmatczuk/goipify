package main

import (
	"net/http"
	ipify "github.com/mmatczuk/goipify"
	"log"
)

func main() {
	log.Println("Starting...")
	if err := http.ListenAndServe(":80", ipify.Ipifi()); err != nil {
		log.Fatalf("Failed to start: %s", err)
	}
	log.Println("Bye")
}
