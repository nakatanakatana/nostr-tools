package main

import (
	"context"
)

type NIP05Provider interface {
	GetPubKey(ctx context.Context, name string) (string, error)
}
