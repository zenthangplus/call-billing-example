package mapper

import (
	"github.com/zenthangplus/call-billing-example/src/adapter/repository/mysql/model"
	"github.com/zenthangplus/call-billing-example/src/core/entity"
	"time"
)

func BillingModelToEntity(m *model.Billing) *entity.Billing {
	return &entity.Billing{
		Id:           m.Id,
		Username:     m.Username,
		CallDuration: time.Duration(m.CallDuration) * time.Millisecond,
		CallCount:    m.CallCount,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
	}
}
