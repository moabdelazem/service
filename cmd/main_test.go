package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestHealthHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := healthHandler(c); err != nil {
		t.Fatalf("Failed to get health status: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}

	// Parse the response body
	var response map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	// Check the values
	if status, ok := response["status"].(string); !ok || status != "OK" {
		t.Errorf("Expected status to be 'OK', got %v", response["status"])
	}

	if code, ok := response["code"].(float64); !ok || int(code) != 200 {
		t.Errorf("Expected code to be 200, got %v", response["code"])
	}
}
