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
