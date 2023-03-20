package port

import (
	"context"
	"github.com/zenthangplus/call-billing-example/src/core/entity"
	"time"
)

type CallRepository interface {
	// Create a call with simple username and call's duration
	Create(ctx context.Context, username string, duration time.Duration) (*entity.Call, error)
}
