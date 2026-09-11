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
		t.Fatalf("期望状态码 200，实际 %d", w.Code)
	}

	var body map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("解析响应失败: %v", err)
	}
	if body["message"] != "hello, gin" {
		t.Fatalf("期望 message 为 \"hello, gin\"，实际 %q", body["message"])
	}
}

func TestHelloDefaultName(t *testing.T) {
	r := router.New()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/hello", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var body map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("解析响应失败: %v", err)
	}
	if body["message"] != "hello, world" {
		t.Fatalf("期望 message 为 \"hello, world\"，实际 %q", body["message"])
	}
}

func TestHealthz(t *testing.T) {
	r := router.New()

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("期望状态码 200，实际 %d", w.Code)
	}
}
