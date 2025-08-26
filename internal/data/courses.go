package data

import "time"

type Course struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
}
