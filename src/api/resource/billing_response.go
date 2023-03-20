package resource

import "github.com/zenthangplus/call-billing-example/src/core/entity"

type BillingResponse struct {
	Id         int     `json:"id"`
	CallCount  int64   `json:"call_count"`
	BlockCount int64   `json:"block_count"`
	Price      float64 `json:"price"`
}

func NewBillingResponse(entity *entity.EstimatedBilling) *BillingResponse {
	return &BillingResponse{
		Id:         entity.Id,
		CallCount:  entity.CallCount,
		BlockCount: entity.BlockCount,
		Price:      entity.Price,
	}
}
