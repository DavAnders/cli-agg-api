package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRespondWithJSON(t *testing.T) {
	w := httptest.NewRecorder()
	testPayload := map[string]string{"test": "payload"}
	expectedBody, _ := json.Marshal(testPayload)

	respondWithJSON(w, http.StatusOK, testPayload)

	// Check the status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Check the response body
	body := w.Body.Bytes()
	if !bytes.Equal(body, expectedBody) {
		t.Errorf("Expected body %s, got %s", expectedBody, body)
	}

	// Check the Content-Type header
	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected Content-Type 'application/json', got '%s'", contentType)
	}
}
