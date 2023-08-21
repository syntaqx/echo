package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEchoHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/echo", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rr := httptest.NewRecorder()

	EchoHandler(rr, req)
	if rr.Body.String() != HelloWorld {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), HelloWorld)
	}
}
