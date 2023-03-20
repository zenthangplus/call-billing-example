package service

import (
	"context"
	"github.com/zenthangplus/call-billing-example/src/core/entity"
)

type BillingService interface {
	// Aggregate billing from a call
	Aggregate(ctx context.Context, call *entity.Call) error

	// Get billing for a username
	Get(ctx context.Context, username string) (*entity.EstimatedBilling, error)
}
