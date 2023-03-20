package resource

import (
	"github.com/zenthangplus/call-billing-example/src/core/entity"
	"reflect"
	"testing"
)

func TestNewBillingResponse(t *testing.T) {
	type args struct {
		entity *entity.EstimatedBilling
	}
	tests := []struct {
		name string
		args args
		want *BillingResponse
	}{
		{
			name: "test full data",
			args: args{entity: &entity.EstimatedBilling{
				Billing:    entity.Billing{Id: 1, CallCount: 3},
				BlockCount: 2,
				Price:      0.3,
			}},
			want: &BillingResponse{
				Id:         1,
				CallCount:  3,
				BlockCount: 2,
				Price:      0.3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBillingResponse(tt.args.entity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBillingResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
