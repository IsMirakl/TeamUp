package model

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID string `gorm:"primaryKey"`

	Title       string   `gorm:"size:50;not null;"`
	Description string   `gorm:"size:750; not null"`
	Tags        []string `gorm:"serializer:json" json:"tags"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	AuthorID string
}
