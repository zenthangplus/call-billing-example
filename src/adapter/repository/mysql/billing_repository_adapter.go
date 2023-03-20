package mysql

import (
	"context"
	"github.com/zenthangplus/call-billing-example/src/adapter/repository/mysql/mapper"
	"github.com/zenthangplus/call-billing-example/src/adapter/repository/mysql/model"
	"github.com/zenthangplus/call-billing-example/src/core/entity"
	"github.com/zenthangplus/call-billing-example/src/core/port"
	"gorm.io/gorm"
	"time"
)

type BillingRepositoryAdapter struct {
	base
}

func NewBillingRepositoryAdapter(db *gorm.DB) *BillingRepositoryAdapter {
	return &BillingRepositoryAdapter{base: base{db: db}}
}

func (b BillingRepositoryAdapter) Begin(ctx context.Context) (port.BillingRepository, error) {
	tx := b.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return NewBillingRepositoryAdapter(tx), nil
}

func (b BillingRepositoryAdapter) Commit(ctx context.Context) error {
	return b.db.WithContext(ctx).Commit().Error
}

func (b BillingRepositoryAdapter) Rollback(ctx context.Context) error {
	return b.db.WithContext(ctx).Rollback().Error
}

func (b BillingRepositoryAdapter) ExistsByUsername(ctx context.Context, username string) (bool, error) {
	var exists bool
	tx := b.db.WithContext(ctx).Model(&model.Billing{}).
		Select("count(*) > 0").
		Where("username = ?", username).
		Find(&exists)
	return exists, tx.Error
}

func (b BillingRepositoryAdapter) Create(ctx context.Context, username string, callDuration time.Duration, callCount int64) error {
	return b.db.WithContext(ctx).
		Create(&model.Billing{
			Username:     username,
			CallDuration: callDuration.Milliseconds(),
			CallCount:    callCount,
		}).Error
}

func (b BillingRepositoryAdapter) IncreaseByUsername(ctx context.Context, username string, increasedDuration time.Duration, increasedCount int64) error {
	return b.db.WithContext(ctx).
		Exec("UPDATE billings SET call_duration = call_duration + ?, call_count = call_count + ? WHERE username = ?",
			increasedDuration.Milliseconds(), increasedCount, username).Error
}

func (b BillingRepositoryAdapter) FindOneByUsername(ctx context.Context, username string) (*entity.Billing, error) {
	var bill = model.Billing{Username: username}
	result := b.db.WithContext(ctx).First(&bill)
	if result.Error != nil {
		return nil, b.handleError(result.Error)
	}
	return mapper.BillingModelToEntity(&bill), nil
}
