package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"<%= modulePath %>/internal/router"
)

func TestHello(t *testing.T) {
	r := router.New()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/hello?name=gin", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var body map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}
	if body["message"] != "hello, gin" {
		t.Fatalf("expected message \"hello, gin\", got %q", body["message"])
	}
}

func TestHelloDefaultName(t *testing.T) {
	r := router.New()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/hello", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var body map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}
	if body["message"] != "hello, world" {
		t.Fatalf("expected message \"hello, world\", got %q", body["message"])
	}
}

func TestHealthz(t *testing.T) {
	r := router.New()

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}
