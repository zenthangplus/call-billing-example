package model

import "time"

type Call struct {
	Id        int `gorm:"primaryKey"`
	Username  string
	Duration  int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
