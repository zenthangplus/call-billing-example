package mapper

import (
	"github.com/zenthangplus/call-billing-example/src/adapter/repository/mysql/model"
	"github.com/zenthangplus/call-billing-example/src/core/entity"
	"reflect"
	"testing"
	"time"
)

func TestBillingModelToEntity(t *testing.T) {
	t1, _ := time.Parse(time.DateTime, "2023-01-02 15:04:05")
	t2, _ := time.Parse(time.DateTime, "2023-01-02 15:04:06")
	type args struct {
		m *model.Billing
	}
	tests := []struct {
		name string
		args args
		want *entity.Billing
	}{
		{
			name: "test full data",
			args: args{
				m: &model.Billing{
					Id:           1,
					Username:     "user1",
					CallDuration: 30000,
					CallCount:    2,
					CreatedAt:    t1,
					UpdatedAt:    t2,
				},
			},
			want: &entity.Billing{
				Id:           1,
				Username:     "user1",
				CallDuration: 30 * time.Second,
				CallCount:    2,
				CreatedAt:    t1,
				UpdatedAt:    t2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BillingModelToEntity(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BillingModelToEntity() = %v, want %v", got, tt.want)
			}
		})
	}
}
