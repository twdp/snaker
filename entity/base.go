package entity

import "time"

type Base struct {
	Id int64

	CreatedAt time.Time

	UpdatedAt time.Time
}