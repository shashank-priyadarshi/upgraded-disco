package common

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

// InvalidEndpoint function hit
// Write to response: Endpoint does not exist
func TestInvalidEndpoint(t *testing.T) {
	// Create a new http GET request
	req, err := http.NewRequest("GET", "/invalid-endpoint", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new recorder
	rr := httptest.NewRecorder()
	// Pass InvalidEndpoint as handler function for the route defined in the http request above
	handler := http.HandlerFunc(InvalidEndpoint)

	var buf bytes.Buffer

	// Serve the request
	log.SetOutput(&buf) // redirect standard output to a buffer
	handler.ServeHTTP(rr, req)
	log.SetOutput(os.Stdout) // restore standard output

	// expectedLog := fmt.Sprintf("Endpoint Hit: %v with %v method\n", req.URL.Path, req.Method)
	// capturedLog := buf.String()

	// if capturedLog != expectedLog {
	// 	t.Errorf("unexpected log output: got %v, want %v", capturedLog, expectedLog)
	// }

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}

	expected := "Endpoint does not exist"
	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
