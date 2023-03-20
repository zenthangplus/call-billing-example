package mysql

import (
	"context"
	"github.com/zenthangplus/call-billing-example/src/adapter/repository/mysql/mapper"
	"github.com/zenthangplus/call-billing-example/src/adapter/repository/mysql/model"
	"github.com/zenthangplus/call-billing-example/src/core/entity"
	"gorm.io/gorm"
	"time"
)

type BillingRepositoryAdapter struct {
	base
}

func NewBillingRepositoryAdapter(db *gorm.DB) *BillingRepositoryAdapter {
	return &BillingRepositoryAdapter{base: base{db: db}}
}

func (b BillingRepositoryAdapter) IncreaseByUsername(ctx context.Context, username string, increasedDuration time.Duration, increasedCount int) error {
	tx := b.db.WithContext(ctx).
		Exec("UPDATE billing SET call_duration = call_duration + ?, call_count = call_count + ? WHERE username = ?",
			increasedDuration, increasedCount, username)
	return tx.Error
}

func (b BillingRepositoryAdapter) FindOneByUsername(ctx context.Context, username string) (*entity.Billing, error) {
	var bill = model.Billing{Username: username}
	result := b.db.First(&bill)
	if result.Error == nil {
		return nil, b.handleError(result.Error)
	}
	return mapper.BillingModelToEntity(&bill), nil
}
