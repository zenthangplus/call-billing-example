package model

import "time"

type Billing struct {
	Id           int `gorm:"primaryKey"`
	Username     string
	CallDuration int64
	CallCount    int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
