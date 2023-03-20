package entity

import (
	"github.com/pkg/errors"
	"time"
)

type EstimatedBilling struct {
	Billing
	BlockTime     time.Duration
	BlockCount    int64
	PricePerBlock float64
	Price         float64
}

func NewEstimatedBilling(billing *Billing, blockTime time.Duration, pricePerBlock float64) (*EstimatedBilling, error) {
	if blockTime <= 0 {
		return nil, errors.New("invalid block time")
	}
	blockCountF := float64(billing.CallDuration) / float64(blockTime)
	blockCount := int64(blockCountF + 0.5)
	price := float64(blockCount) * pricePerBlock
	return &EstimatedBilling{
		Billing:       *billing,
		BlockTime:     blockTime,
		BlockCount:    blockCount,
		PricePerBlock: pricePerBlock,
		Price:         price,
	}, nil
}
