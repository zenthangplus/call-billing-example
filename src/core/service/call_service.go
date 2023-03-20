package service

import (
	"context"
	"github.com/zenthangplus/call-billing-example/src/core/entity"
)

type CallService interface {
	// EndCall ends a call
	EndCall(ctx context.Context, username string, durationMs int64) (*entity.Call, error)
}
