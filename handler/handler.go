package handler

import (
	"fmt"
	"net/http"
)

const HelloWorld = "Hello, world!"

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, HelloWorld)
}
