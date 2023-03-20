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

func NewDefaultBillingService(
	repo port.BillingRepository,
	conf *config.BillingConfig,
) *DefaultBillingService {
	return &DefaultBillingService{
		repo: repo,
		conf: conf,
	}
}

func (d DefaultBillingService) Aggregate(ctx context.Context, call *entity.Call) error {
	txRepo, err := d.repo.Begin(ctx)
	if err != nil {
		return errors.WithMessage(err, "Cannot start transaction")
	}
	defer txRepo.Rollback(ctx)

	exists, err := txRepo.ExistsByUsername(ctx, call.Username)
	if err != nil {
		return errors.WithMessage(err, "cannot check billing is exists or not")
	}
	if !exists {
		if err := txRepo.Create(ctx, call.Username, call.Duration, 1); err != nil {
			return errors.WithMessage(err, "create billing failed")
		}
		if err := txRepo.Commit(ctx); err != nil {
			return errors.WithMessage(err, "create billing failed")
		}
		return nil
	}
	if err := d.repo.IncreaseByUsername(ctx, call.Username, call.Duration, 1); err != nil {
		return errors.WithMessage(err, "cannot increase billing data")
	}
	if err := txRepo.Commit(ctx); err != nil {
		return errors.WithMessage(err, "create billing failed")
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
