package models

import "time"

type User struct {
	ID           int
	Phone        string
	Password     string
	RegisteredAt time.Time
	ReferredID   int
	StatusID     int
	CreateAt     time.Time
	UpdatedAt    time.Time
}
