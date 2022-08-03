package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	srv := &http.Server{
		Addr:    net.JoinHostPort("", port),
		Handler: http.HandlerFunc(handler),
	}

	log.Printf("http server listening at %s\n", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}
