package service

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zenthangplus/call-billing-example/src/core/entity"
	"github.com/zenthangplus/call-billing-example/src/core/event"
	"github.com/zenthangplus/call-billing-example/src/core/port"
	"time"
)

type DefaultCallService struct {
	repo      port.CallRepository
	publisher port.EventPublisher
}

func NewDefaultCallService(
	repo port.CallRepository,
	publisher port.EventPublisher,
) *DefaultCallService {
	return &DefaultCallService{
		repo:      repo,
		publisher: publisher,
	}
}

func (d DefaultCallService) EndCall(ctx context.Context, username string, durationMs int64) (*entity.Call, error) {
	if username == "" {
		return nil, errors.New("invalid username")
	}
	if durationMs <= 0 {
		return nil, errors.New("invalid duration")
	}
	call, err := d.repo.Create(ctx, username, time.Duration(durationMs)*time.Millisecond)
	if err != nil {
		return nil, errors.WithMessage(err, "create call error")
	}
	d.publisher.Publish(event.NewCallEnded(ctx, call))
	return call, nil
}
