package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"

	"golang.org/x/net/http2"
)

func main() {
	addr := flag.String("addr", "http://localhost:8765", "Address of server")
	flag.Parse()

	transport := &http2.Transport{
		AllowHTTP: true,
		// Pretend we are dialing a TLS endpoint.
		// Note, we ignore the passed tls.Config
		DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
			return net.Dial(network, addr)
		},
	}

	client := http.Client{Transport: transport}

	req, err := http.NewRequest(http.MethodGet, *addr, nil)
	if err != nil {
		log.Fatalf("NewRequest: %s\n", err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("get error: %s\n", err)
	}

	_, _ = io.Copy(os.Stdout, res.Body)
	_ = res.Body.Close()
	fmt.Println()
	for k, v := range res.Trailer {
		fmt.Printf("%s: %v\n", k, v)
	}
}
