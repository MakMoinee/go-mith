package idemstore

import "errors"

var (
	ErrInProgress = errors.New("idempotent request is already being processed")
	ErrConflict   = errors.New("idempotency key is already associated with a different request")
)
