package service

import (
	"context"
	"github.com/zenthangplus/call-billing-example/src/core/entity"
)

type BillingService interface {
	Aggregate(ctx context.Context, call *entity.Call) error
	Get(ctx context.Context, username string) (*entity.EstimatedBilling, error)
}
