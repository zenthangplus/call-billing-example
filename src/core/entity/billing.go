package entity

import "time"

type Billing struct {
	Id           int
	Username     string
	CallDuration time.Duration
	CallCount    int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
