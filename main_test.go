package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	go main()
	os.Exit(m.Run())
}

func TestHandler(t *testing.T) {
	req, err := http.NewRequest(
		http.MethodGet,
		"http://localhost:8080/weather",
		nil,
	)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rec := httptest.NewRecorder()
	handler(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("Expected code 200, got %d", rec.Code)
	}
}
