package main

import (
	"context"
	"errors"
	"strings"
)

type NIP05Provider interface {
	GetPubKey(ctx context.Context, name string) (string, error)
	GetRelays(ctx context.Context, pubkey string) ([]string, error)
}

type MemoryProvider struct {
	mapping map[string]string
	relays  map[string][]string
}

func NewMemoryProvider(mapping map[string]string, relaysConfig map[string]string) *MemoryProvider {
	relays := make(map[string][]string)
	for k, v := range relaysConfig {
		// Split comma-separated relays
		list := strings.Split(v, ",")
		for i := range list {
			list[i] = strings.TrimSpace(list[i])
		}
		relays[k] = list
	}

	return &MemoryProvider{
		mapping: mapping,
		relays:  relays,
	}
}

func (m *MemoryProvider) GetPubKey(ctx context.Context, name string) (string, error) {
	pubkey, ok := m.mapping[name]
	if !ok {
		return "", errors.New("name not found")
	}
	return pubkey, nil
}

func (m *MemoryProvider) GetRelays(ctx context.Context, pubkey string) ([]string, error) {
	r, ok := m.relays[pubkey]
	if !ok {
		return nil, nil
	}
	return r, nil
}