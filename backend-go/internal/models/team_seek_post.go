package models

type TeamSeekPost struct {
	ID string `gorm:"primaryKey"`
	Title string `gorm:"size:50;not null;"`
	Description string `gorm:"size:750; not null"`
	Tags []string
	Author Author `gorm:"foreignKey:ID"`
}

type Author struct {
	ID string `gorm:"primaryKey"`
	User User `gorm:"foreignKey:UserID"`
}