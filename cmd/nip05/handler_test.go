package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNIP05Handler(t *testing.T) {
	mapping := map[string]string{
		"bob":   "pubkey1",
		"alice": "pubkey2",
	}
	provider := NewMemoryProvider(mapping)
	handler := NewNIP05Handler(provider)

	tests := []struct {
		name           string
		queryName      string
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name:           "Valid Name Bob",
			queryName:      "bob",
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"names": map[string]interface{}{
					"bob": "pubkey1",
				},
			},
		},
		{
			name:           "Valid Name Alice",
			queryName:      "alice",
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"names": map[string]interface{}{
					"alice": "pubkey2",
				},
			},
		},
		{
			name:           "Unknown Name",
			queryName:      "unknown",
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"names": map[string]interface{}{},
			},
		},
		{
			name:           "Missing Name Param",
			queryName:      "",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/.well-known/nostr.json?name="+tt.queryName, nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()

			// We need to serve via the handler.
			// If NewNIP05Handler returns a http.Handler, we use it.
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}

			if tt.expectedBody != nil {
				var gotBody map[string]interface{}
				if err := json.Unmarshal(rr.Body.Bytes(), &gotBody); err != nil {
					t.Fatalf("Failed to decode response body: %v", err)
				}
				
				// Verify structure (simplistic check)
				gotNames, ok := gotBody["names"].(map[string]interface{})
				if !ok {
					t.Errorf("Response body missing 'names' object")
					return
				}
				
				expectedNames := tt.expectedBody["names"].(map[string]interface{})
				if len(gotNames) != len(expectedNames) {
					t.Errorf("Expected %d names, got %d", len(expectedNames), len(gotNames))
				}
				
				for k, v := range expectedNames {
					if gotNames[k] != v {
						t.Errorf("Expected mapping %s -> %s, got %s", k, v, gotNames[k])
					}
				}
			}
		})
	}
}

func TestNIP05Handler_CORS(t *testing.T) {
	mapping := map[string]string{}
	provider := NewMemoryProvider(mapping)
	handler := NewNIP05Handler(provider)
	
	req, err := http.NewRequest("GET", "/.well-known/nostr.json?name=bob", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	// Apply Middleware
	corsHandler := CORSMiddleware(handler)
	corsHandler.ServeHTTP(rr, req)

	if val := rr.Header().Get("Access-Control-Allow-Origin"); val != "*" {
		t.Errorf("Expected Access-Control-Allow-Origin: *, got %s", val)
	}
}
