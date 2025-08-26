package data

import (
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
}
