package idemstore

import (
	"context"
	"time"
)

type Store interface {
	Execute(
		ctx context.Context,
		key string,
		requestHash string,
		fn func(context.Context) ([]byte, error),
	) ([]byte, error)
}

type Config struct {
	TTL           time.Duration
	ProcessingTTL time.Duration
	PollInterval  time.Duration
	WaitTimeout   time.Duration
}

func DefaultConfig() Config {
	return Config{
		TTL:           24 * time.Hour,
		ProcessingTTL: 5 * time.Minute,
		PollInterval:  100 * time.Millisecond,
		WaitTimeout:   10 * time.Second,
	}
}
