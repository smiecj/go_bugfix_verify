package main

import (
	"log"
	"net/http"
)

// https://github.com/golang/go/issues/47143
// crypto/tls: clients can panic when provided a certificate of the wrong type for the negotiated parameters (CVE-2021-34558)
func main() {
	// todo: use bad certification to connect https website
	_, err := http.Get("https://gocn.vip/")
	if err != nil {
		log.Fatal(err)
	}
	log.Print("connect success")
}
