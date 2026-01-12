package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	if err := Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func Run() error {
	// 1. Load Configuration
	cfg, err := LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	// 2. Setup Logger
	logger := SetupLogger(cfg.LogLevel)
	slog.SetDefault(logger)
	logger.Info("Starting NIP-05 Server", "port", cfg.Port, "domain", cfg.Domain)

	// 3. Initialize Data Provider
	provider := NewMemoryProvider(cfg.Mapping)

	// 4. Initialize Handler
	handler := NewNIP05Handler(provider)
	router := CORSMiddleware(handler)

	// 5. Setup Server Mux
	mux := http.NewServeMux()
	mux.Handle("/.well-known/nostr.json", router)

	// 6. Start Server
	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	logger.Info("Listening on", "addr", addr)
	return http.ListenAndServe(addr, mux)
}