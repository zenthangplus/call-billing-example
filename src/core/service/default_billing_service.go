package service

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zenthangplus/call-billing-example/src/core/config"
	"github.com/zenthangplus/call-billing-example/src/core/entity"
	"github.com/zenthangplus/call-billing-example/src/core/port"
)

type DefaultBillingService struct {
	repo port.BillingRepository
	conf *config.BillingConfig
}

func (d DefaultBillingService) Aggregate(ctx context.Context, call *entity.Call) error {
	err := d.repo.IncreaseByUsername(ctx, call.Username, call.Duration, 1)
	if err != nil {
		return errors.WithMessage(err, "cannot increase billing data")
	}
	return nil
}

func (d DefaultBillingService) Get(ctx context.Context, username string) (*entity.EstimatedBilling, error) {
	billing, err := d.repo.FindOneByUsername(ctx, username)
	if err != nil {
		return nil, errors.WithMessage(err, "cannot find billing for user")
	}
	return entity.NewEstimatedBilling(billing, d.conf.BlockTime, d.conf.PricePerBlock)
}
