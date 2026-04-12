package model

import (
	"time"
)

type Post struct {
	ID string

	Title       string
	Description string
	Tags        []string `json:"tags"`

	CreatedAt time.Time
	UpdatedAt time.Time

	AuthorID string
}
