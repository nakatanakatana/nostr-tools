package main

import (
	"context"
	"errors"
	"os"
	"reflect"
	"testing"
)

type MockProvider struct {
	mapping map[string]string
	relays  map[string][]string
}

func (m *MockProvider) GetPubKey(ctx context.Context, name string) (string, error) {
	pubkey, ok := m.mapping[name]
	if !ok {
		return "", errors.New("not found")
	}
	return pubkey, nil
}

func (m *MockProvider) GetRelays(ctx context.Context, pubkey string) ([]string, error) {
	relays, ok := m.relays[pubkey]
	if !ok {
		return nil, nil
	}
	return relays, nil
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
	relays := map[string]string{"pubkey1": "wss://r1,wss://r2"}
	provider := NewMemoryProvider(mapping, relays)

	pubkey, err := provider.GetPubKey(context.Background(), "bob")
	if err != nil {
		t.Fatalf("Expected nil error, got %v", err)
	}
	if pubkey != "pubkey1" {
		t.Errorf("Expected pubkey1, got %s", pubkey)
	}

	r, err := provider.GetRelays(context.Background(), "pubkey1")
	if err != nil {
		t.Fatalf("Expected nil error for relays, got %v", err)
	}
	if len(r) != 2 {
		t.Errorf("Expected 2 relays, got %d", len(r))
	}
	if r[0] != "wss://r1" || r[1] != "wss://r2" {
		t.Errorf("Relays mismatch: %v", r)
	}

	_, err = provider.GetPubKey(context.Background(), "alice")
	if err == nil {
		t.Error("Expected error for missing name, got nil")
	}
}

func TestParseRelays(t *testing.T) {
	input := map[string]string{
		"pub1": "wss://r1, wss://r2",
		"pub2": "wss://r3",
	}
	expected := map[string][]string{
		"pub1": {"wss://r1", "wss://r2"},
		"pub2": {"wss://r3"},
	}
	got := ParseRelays(input)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ParseRelays() = %v, want %v", got, expected)
	}
}

func TestConfigToMemoryProvider(t *testing.T) {
	if err := os.Setenv("NIP05_MAPPING", "user1:pub1,user2:pub2"); err != nil {
		t.Fatal(err)
	}
	if err := os.Setenv("NIP05_RELAYS", "pub1:wss://r1,wss://r2|pub2:wss://r3"); err != nil {
		t.Fatal(err)
	}
	defer os.Clearenv()

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	provider := NewMemoryProvider(cfg.Mapping, cfg.Relays)

	tests := []struct {
		name       string
		wantPubKey string
		wantRelays []string
		wantErr    bool
	}{
		{"user1", "pub1", []string{"wss://r1", "wss://r2"}, false},
		{"user2", "pub2", []string{"wss://r3"}, false},
		{"unknown", "", nil, true},
	}

	for _, tt := range tests {
		got, err := provider.GetPubKey(context.Background(), tt.name)
		if (err != nil) != tt.wantErr {
			t.Errorf("GetPubKey(%s) error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.wantPubKey {
			t.Errorf("GetPubKey(%s) got = %s, want %s", tt.name, got, tt.wantPubKey)
		}

		if !tt.wantErr {
			gotRelays, err := provider.GetRelays(context.Background(), got)
			if err != nil {
				t.Errorf("GetRelays(%s) error = %v", got, err)
			}
			if !reflect.DeepEqual(gotRelays, tt.wantRelays) {
				t.Errorf("GetRelays(%s) got = %v, want %v", got, gotRelays, tt.wantRelays)
			}
		}
	}
}
