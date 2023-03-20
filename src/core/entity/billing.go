package entity

import "time"

type Billing struct {
	Id           int
	Username     string
	CallDuration time.Duration
	CallCount    int64
	BlockCount   int64
	Price        float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
