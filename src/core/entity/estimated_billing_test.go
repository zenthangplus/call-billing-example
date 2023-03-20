package entity

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestNewEstimatedBilling(t *testing.T) {
	t1 := time.Unix(1679330938, 0)
	t2 := time.Unix(1679330940, 0)
	type args struct {
		billing       *Billing
		blockTime     time.Duration
		pricePerBlock float64
	}
	tests := []struct {
		name    string
		args    args
		want    *EstimatedBilling
		wantErr bool
	}{
		{
			name: "when invalid block time should error",
			args: args{
				billing: &Billing{
					Id:           1,
					Username:     "user1",
					CallDuration: 70 * time.Second,
					CallCount:    3,
					CreatedAt:    t1,
					UpdatedAt:    t2,
				},
				blockTime:     0,
				pricePerBlock: 0.2,
			},
			wantErr: true,
			want:    nil,
		},
		{
			name: "call duration upper block time",
			args: args{
				billing: &Billing{
					Id:           1,
					Username:     "user1",
					CallDuration: 70 * time.Second,
					CallCount:    3,
					CreatedAt:    t1,
					UpdatedAt:    t2,
				},
				blockTime:     30 * time.Second,
				pricePerBlock: 0.2,
			},
			want: &EstimatedBilling{
				Billing: Billing{
					Id:           1,
					Username:     "user1",
					CallDuration: 70 * time.Second,
					CallCount:    3,
					CreatedAt:    t1,
					UpdatedAt:    t2,
				},
				BlockTime:     30 * time.Second,
				BlockCount:    3,
				PricePerBlock: 0.2,
				Price:         0.6,
			},
		},
		{
			name: "call duration equal block time",
			args: args{
				billing: &Billing{
					Id:           1,
					Username:     "user1",
					CallDuration: 30 * time.Second,
					CallCount:    3,
					CreatedAt:    t1,
					UpdatedAt:    t2,
				},
				blockTime:     30 * time.Second,
				pricePerBlock: 0.2,
			},
			want: &EstimatedBilling{
				Billing: Billing{
					Id:           1,
					Username:     "user1",
					CallDuration: 30 * time.Second,
					CallCount:    3,
					CreatedAt:    t1,
					UpdatedAt:    t2,
				},
				BlockTime:     30 * time.Second,
				BlockCount:    1,
				PricePerBlock: 0.2,
				Price:         0.2,
			},
		},
		{
			name: "call duration under block time",
			args: args{
				billing: &Billing{
					Id:           1,
					Username:     "user1",
					CallDuration: 1 * time.Second,
					CallCount:    3,
					CreatedAt:    t1,
					UpdatedAt:    t2,
				},
				blockTime:     30 * time.Second,
				pricePerBlock: 0.2,
			},
			want: &EstimatedBilling{
				Billing: Billing{
					Id:           1,
					Username:     "user1",
					CallDuration: 1 * time.Second,
					CallCount:    3,
					CreatedAt:    t1,
					UpdatedAt:    t2,
				},
				BlockTime:     30 * time.Second,
				BlockCount:    1,
				PricePerBlock: 0.2,
				Price:         0.2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEstimatedBilling(tt.args.billing, tt.args.blockTime, tt.args.pricePerBlock)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewEstimatedBilling() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil {
				if !assert.InDelta(t, tt.want.Price, got.Price, 10e-9) {
					t.Errorf("NewEstimatedBilling() got.price = %v, want.price %v", got.Price, tt.want.Price)
					return
				}
				tt.want.Price = got.Price // don't want to assert anymore price due by we asserted above
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEstimatedBilling() got = %v, want %v", got, tt.want)
			}
		})
	}
}
