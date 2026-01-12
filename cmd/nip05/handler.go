package main

import (
	"encoding/json"
	"net/http"
)

type NIP05Handler struct {
	provider NIP05Provider
}

func NewNIP05Handler(provider NIP05Provider) *NIP05Handler {
	return &NIP05Handler{
		provider: provider,
	}
}

func (h *NIP05Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "missing name parameter", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// CORS header will be added by middleware, but good to ensure response is valid JSON.

	response := map[string]interface{}{
		"names":  map[string]string{},
		"relays": map[string][]string{},
	}

	pubkey, err := h.provider.GetPubKey(r.Context(), name)
	if err == nil {
		// Found
		response["names"].(map[string]string)[name] = pubkey
	} else {
		// Not found, return empty names map (or just don't add it)
		// NIP-05: "If the identifier is not found, the server SHOULD return 404 Not Found" is NOT in NIP-05.
		// NIP-05 says: "The response should be a JSON object..."
		// Common practice: Return 200 with empty JSON or the requested name missing.
		// However, for single name query, if it's missing, often it is just empty names.
		// Let's stick to returning 200 OK with empty result if not found, as per test expectation.
		// Wait, test expects 404 for unknown? 
		// "expectedStatus: http.StatusNotFound" in test case "Unknown Name".
		// Let's adjust implementation to match test expectation or standard.
		// Standard NIP-05 doesn't explicitly mandate 404. But 200 OK with empty is safer for clients.
		// I will update the test to expect 200 OK for unknown name, as that is more robust.
		// But for now, let's implement what seems logical: 
		// If explicit name requested and not found -> empty result (200 OK).
	}

	// Wait, I put expectedStatus: http.StatusNotFound in the test for "Unknown Name". 
	// I should probably return 200 OK with empty map, as that is valid JSON.
	// Returning 404 might be interpreted as "endpoint not found" by some clients.
	// I will update the test to expect 200 OK.
	
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
