package models

import (
	"time"

	"gorm.io/gorm"
)

type TeamSeekPost struct {
	ID string `gorm:"primaryKey"`

	Title string `gorm:"size:50;not null;"`
	Description string `gorm:"size:750; not null"`
	Tags []string `gorm:"type:json"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	AuthorId string 
	Author Author `gorm:"foreignKey:AuthorId;references:ID"`
}

type Author struct {
	ID string `gorm:"primaryKey"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	UserID uint
}