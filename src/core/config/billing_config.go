package config

import "time"

type BillingConfig struct {
	BlockTime     time.Duration
	PricePerBlock float64
}
