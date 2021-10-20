package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	addr := flag.String("bind", ":8765", "Bind address")
	flag.Parse()

	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("proto=%s\n", req.Proto)

		for k, v := range req.Header {
			fmt.Printf("%s: %v\n", k, v)
		}

		start := time.Now()

		w.Header().Add("Trailer", "Timing")
		w.WriteHeader(http.StatusOK)

		_, _ = w.Write([]byte("Hello, world!"))

		w.Header().Set("Timing", strconv.FormatFloat(time.Now().Sub(start).Seconds(), 'f', -1, 64))
	})

	h2s := &http2.Server{}

	server := &http.Server{
		Addr:    *addr,
		Handler: h2c.NewHandler(handler, h2s),
	}
	_ = server.ListenAndServe()
}
