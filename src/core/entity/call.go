package entity

import "time"

type Call struct {
	Id        int
	Username  string
	Duration  time.Duration
	CreatedAt time.Time
	UpdatedAt time.Time
}
