package main

import (
	"net/http"
	ipify "github.com/mmatczuk/goipify"
)

func main() {
	http.ListenAndServe(":80", ipify.Ipifi())
}
