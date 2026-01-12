package main

import (
	"context"
	"errors"
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