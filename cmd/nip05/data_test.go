package main

import (
	"context"
	"errors"
	"os"
	"testing"
)

type MockProvider struct {
	mapping map[string]string
}

func (m *MockProvider) GetPubKey(ctx context.Context, name string) (string, error) {
	pubkey, ok := m.mapping[name]
	if !ok {
		return "", errors.New("not found")
	}
	return pubkey, nil
}

func TestDataInterface(t *testing.T) {
	var provider NIP05Provider = &MockProvider{
		mapping: map[string]string{"bob": "pubkey1"},
	}

	pubkey, err := provider.GetPubKey(context.Background(), "bob")
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
	if pubkey != "pubkey1" {
		t.Errorf("Expected pubkey1, got %s", pubkey)
	}

	_, err = provider.GetPubKey(context.Background(), "alice")
	if err == nil {
		t.Error("Expected error for missing key, got nil")
	}
}

func TestMemoryProvider(t *testing.T) {
	mapping := map[string]string{"bob": "pubkey1"}
	provider := NewMemoryProvider(mapping)

	pubkey, err := provider.GetPubKey(context.Background(), "bob")
	if err != nil {
		t.Fatalf("Expected nil error, got %v", err)
	}
	if pubkey != "pubkey1" {
		t.Errorf("Expected pubkey1, got %s", pubkey)
	}

	_, err = provider.GetPubKey(context.Background(), "alice")
	if err == nil {
		t.Error("Expected error for missing name, got nil")
	}
}

func TestConfigToMemoryProvider(t *testing.T) {
	if err := os.Setenv("NIP05_DOMAIN", "example.com"); err != nil {
		t.Fatal(err)
	}
	if err := os.Setenv("NIP05_MAPPING", "user1:pub1,user2:pub2"); err != nil {
		t.Fatal(err)
	}
	defer os.Clearenv()

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	provider := NewMemoryProvider(cfg.Mapping)

	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{"user1", "pub1", false},
		{"user2", "pub2", false},
		{"unknown", "", true},
	}

	for _, tt := range tests {
		got, err := provider.GetPubKey(context.Background(), tt.name)
		if (err != nil) != tt.wantErr {
			t.Errorf("GetPubKey(%s) error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("GetPubKey(%s) got = %s, want %s", tt.name, got, tt.want)
		}
	}
}