package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestFullIntegration(t *testing.T) {
	// 1. Setup Environment
	os.Clearenv()
	os.Setenv("NIP05_MAPPING", "integration:hexpubkey")
	defer os.Clearenv()

	// 2. Load Config
	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	// 3. Initialize Components
	provider := NewMemoryProvider(cfg.Mapping)
	handler := NewNIP05Handler(provider)
	router := CORSMiddleware(handler)

	// 4. Perform Request
	req, err := http.NewRequest("GET", "/.well-known/nostr.json?name=integration", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	// 5. Verify Response
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}

	// Check CORS
	if val := rr.Header().Get("Access-Control-Allow-Origin"); val != "*" {
		t.Errorf("Expected CORS header *, got %s", val)
	}

	// Check JSON Body
	var body map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &body); err != nil {
		t.Fatalf("Failed to decode body: %v", err)
	}

	names, ok := body["names"].(map[string]interface{})
	if !ok {
		t.Fatal("Response missing 'names' object")
	}

	if names["integration"] != "hexpubkey" {
		t.Errorf("Expected integration->hexpubkey, got %v", names["integration"])
	}
}
