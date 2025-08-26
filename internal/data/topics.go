package data

import (
	"time"
)

type Topic struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	Status    string    `json:"status"`
	AuthorID  int64     `json:"authorId"`
	CourseID  int64     `json:"courseId"`
	Replies   []Reply   `json:"replies"`
}
