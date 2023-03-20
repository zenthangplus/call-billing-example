package mysql

import (
	"context"
	"github.com/zenthangplus/call-billing-example/src/adapter/repository/mysql/mapper"
	"github.com/zenthangplus/call-billing-example/src/adapter/repository/mysql/model"
	"github.com/zenthangplus/call-billing-example/src/core/entity"
	"gorm.io/gorm"
	"time"
)

type CallRepositoryAdapter struct {
	base
}

func NewCallRepositoryAdapter(db *gorm.DB) *CallRepositoryAdapter {
	return &CallRepositoryAdapter{base: base{db: db}}
}

func (c CallRepositoryAdapter) Create(ctx context.Context, username string, duration time.Duration) (*entity.Call, error) {
	m := model.Call{
		Username: username,
		Duration: duration.Milliseconds(),
	}
	c.db.Create(&m)
	return mapper.CallModelToEntity(&m), nil
}
