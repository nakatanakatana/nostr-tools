package main

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Set environment variables
	if err := os.Setenv("NIP05_PORT", "9090"); err != nil {
		t.Fatal(err)
	}
	if err := os.Setenv("NIP05_HOST", "127.0.0.1"); err != nil {
		t.Fatal(err)
	}
	if err := os.Setenv("NIP05_DOMAIN", "example.com"); err != nil {
		t.Fatal(err)
	}
	if err := os.Setenv("NIP05_MAPPING", "bob:pubkey1,alice:pubkey2"); err != nil {
		t.Fatal(err)
	}
	if err := os.Setenv("LOG_LEVEL", "debug"); err != nil {
		t.Fatal(err)
	}
	// Clean up after test
	defer func() {
		_ = os.Unsetenv("NIP05_PORT")
		_ = os.Unsetenv("NIP05_HOST")
		_ = os.Unsetenv("NIP05_DOMAIN")
		_ = os.Unsetenv("NIP05_MAPPING")
		_ = os.Unsetenv("LOG_LEVEL")
	}()

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	if cfg.Port != "9090" {
		t.Errorf("Expected Port 9090, got %s", cfg.Port)
	}
	if cfg.Host != "127.0.0.1" {
		t.Errorf("Expected Host 127.0.0.1, got %s", cfg.Host)
	}
	if cfg.Domain != "example.com" {
		t.Errorf("Expected Domain example.com, got %s", cfg.Domain)
	}
	if cfg.LogLevel != "debug" {
		t.Errorf("Expected LogLevel debug, got %s", cfg.LogLevel)
	}

	// Check Map
	if len(cfg.Mapping) != 2 {
		t.Errorf("Expected 2 mappings, got %d", len(cfg.Mapping))
	}
	if cfg.Mapping["bob"] != "pubkey1" {
		t.Errorf("Expected bob->pubkey1, got %s", cfg.Mapping["bob"])
	}
	if cfg.Mapping["alice"] != "pubkey2" {
		t.Errorf("Expected alice->pubkey2, got %s", cfg.Mapping["alice"])
	}
}

func TestLoadConfigDefaults(t *testing.T) {
	// Ensure cleanup
	os.Clearenv()
	// Set required domain
	if err := os.Setenv("NIP05_DOMAIN", "example.com"); err != nil {
		t.Fatal(err)
	}
	defer func() { _ = os.Unsetenv("NIP05_DOMAIN") }()

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	if cfg.Port != "8080" {
		t.Errorf("Expected default Port 8080, got %s", cfg.Port)
	}
	if cfg.Host != "0.0.0.0" {
		t.Errorf("Expected default Host 0.0.0.0, got %s", cfg.Host)
	}
}

func TestLoadConfig_MissingRequired(t *testing.T) {
	os.Clearenv()
	// Domain is missing
	if err := os.Setenv("NIP05_PORT", "8080"); err != nil {
		t.Fatal(err)
	}
	defer func() { _ = os.Unsetenv("NIP05_PORT") }()

	_, err := LoadConfig()
	if err == nil {
		t.Fatal("Expected error due to missing required field, got nil")
	}
}
