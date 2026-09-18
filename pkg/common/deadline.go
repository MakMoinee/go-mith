package common

import (
	"context"
	"time"
)

func ContextTimeout(ctx context.Context) (time.Duration, bool) {
	deadline, ok := ctx.Deadline()
	if !ok {
		return 0, false
	}

	return time.Until(deadline), true
}
