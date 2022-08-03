package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	srv := &http.Server{
		Addr:    net.JoinHostPort("", port),
		Handler: h2c.NewHandler(http.HandlerFunc(handler), &http2.Server{}),
	}

	log.Printf("http server listening at %s\n", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
