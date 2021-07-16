package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
)

// stack: https://stackoverflow.com/questions/12122159/how-to-do-a-https-request-with-bad-certificate
// https://github.com/golang/go/issues/47143
// crypto/tls: clients can panic when provided a certificate of the wrong type for the negotiated parameters (CVE-2021-34558)
func main() {
	// todo: use bad certification to connect https website
	caCert, err := ioutil.ReadFile("golang.crt")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
				//InsecureSkipVerify: true,
			},
		},
	}

	_, err = client.Get("https://gocn.vip/")
	if err != nil {
		log.Fatal(err)
	}
	log.Print("connect success")
}
