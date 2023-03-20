package mapper

import (
	"github.com/zenthangplus/call-billing-example/src/adapter/repository/mysql/model"
	"github.com/zenthangplus/call-billing-example/src/core/entity"
	"time"
)

func CallModelToEntity(m *model.Call) *entity.Call {
	return &entity.Call{
		Id:        m.Id,
		Username:  m.Username,
		Duration:  time.Duration(m.Duration) * time.Millisecond,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
