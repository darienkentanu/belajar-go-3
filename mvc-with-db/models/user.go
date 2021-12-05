package models

import "time"

type User struct {
	ID        int64
	Name      string
	Email     string
	Password  []byte
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
