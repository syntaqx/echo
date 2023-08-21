package handler

import (
	"fmt"
	"net/http"
)

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}
