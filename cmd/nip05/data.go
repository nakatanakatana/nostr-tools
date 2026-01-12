package main

import (
	"context"
	"errors"
)

type NIP05Provider interface {
	GetPubKey(ctx context.Context, name string) (string, error)
}

type MemoryProvider struct {
	mapping map[string]string
}

func NewMemoryProvider(mapping map[string]string) *MemoryProvider {
	return &MemoryProvider{
		mapping: mapping,
	}
}

func (m *MemoryProvider) GetPubKey(ctx context.Context, name string) (string, error) {
	pubkey, ok := m.mapping[name]
	if !ok {
		return "", errors.New("name not found")
	}
	return pubkey, nil
}
