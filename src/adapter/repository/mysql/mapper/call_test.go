package mapper

import (
	"github.com/zenthangplus/call-billing-example/src/adapter/repository/mysql/model"
	"github.com/zenthangplus/call-billing-example/src/core/entity"
	"reflect"
	"testing"
	"time"
)

func TestCallModelToEntity(t *testing.T) {
	t1, _ := time.Parse(time.DateTime, "2023-01-02 15:04:05")
	t2, _ := time.Parse(time.DateTime, "2023-01-02 15:04:06")
	type args struct {
		m *model.Call
	}
	tests := []struct {
		name string
		args args
		want *entity.Call
	}{
		{
			name: "test full data",
			args: args{m: &model.Call{
				Id:        1,
				Username:  "user1",
				Duration:  20000,
				CreatedAt: t1,
				UpdatedAt: t2,
			}},
			want: &entity.Call{
				Id:        1,
				Username:  "user1",
				Duration:  20 * time.Second,
				CreatedAt: t1,
				UpdatedAt: t2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CallModelToEntity(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CallModelToEntity() = %v, want %v", got, tt.want)
			}
		})
	}
}
