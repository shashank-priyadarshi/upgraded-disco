package common

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBearerAuthAPICall(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected %v request, got %v", http.MethodGet, r.Method)
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("expected Content-Type header to be %v, got %v", "application/json", r.Header.Get("Content-Type"))
		}
		if r.Header.Get("Authorization") != "Bearer abc123" {
			t.Errorf("expected Authorization header to be %v, got %v", "Bearer abc123", r.Header.Get("Authorization"))
		}
		fmt.Fprintln(w, "test response")
	}))
	defer ts.Close()

	resp, status := BearerAuthAPICall(ts.URL, "abc123")

	if status != http.StatusOK {
		t.Errorf("expected status %v, got %v", http.StatusOK, status)
	}

	expected := []byte("test response")
	if !strings.EqualFold(strings.TrimSpace(string(resp)), strings.TrimSpace(string(expected))) {
		t.Errorf("expected response %v, got %v", expected, resp)
	}
}

func TestNoAuthAPICall(t *testing.T) {
	// Create a test server to simulate the API endpoint
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Unsupported media type", http.StatusUnsupportedMediaType)
			return
		}
		if r.Header.Get("Origin") != "test-origin" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"success": true}`)
	}))
	defer ts.Close()

	// Make a request to the test server using the NoAuthAPICall function
	reqBody := []byte(strings.TrimSpace(`{"data": "test"}`))
	respBody, statusCode := NoAuthAPICall(ts.URL, "test-origin", reqBody)

	// Check the response body and status code
	expectedBody := []byte(strings.TrimSpace(`{"success": true}`))
	if !strings.EqualFold(strings.TrimSpace(string(respBody)), strings.TrimSpace(string(expectedBody))) {
		t.Errorf("unexpected response body: got %v want %v", string(respBody), string(expectedBody))
	}
	if statusCode != http.StatusOK {
		t.Errorf("unexpected status code: got %v want %v", statusCode, http.StatusOK)
	}
}
