package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEchoHandler(t *testing.T) {
	// 1. Create a new request
	req, err := http.NewRequest("GET", "/echo", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	// 2. Create a response recorder
	rr := httptest.NewRecorder()

	// 3. Call the EchoHandler function
	EchoHandler(rr, req)

	// 4. Check the recorded response
	expected := "Hello, world!"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
