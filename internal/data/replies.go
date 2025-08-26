package data

import (
	"time"
)

type Reply struct {
	ID         int64     `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	Message    string    `json:"message"`
	TopicID    int64     `json:"topicId"`
	AuthorID   int64     `json:"authorId"`
	IsSolution bool      `json:"isSolution"`
}
