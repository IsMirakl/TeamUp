package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	UserID string	`gorm:"primaryKey"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Email string `gorm:"unique;not null"`
	EmailVerified bool

	Account *Account `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE"`

	Name string `gorm:"size:25;not null;"`
	Avatar *string

	Role Role	`gorm:"type:varchar(20);default:'user'"`
	SubscriptionPlan SubscriptionPlan `gorm:"type:varchar(20);default:'Free'"`
}

type Account struct {
	UserID string `gorm:"primaryKey"`

	PasswordHash  string `gorm:"size:255"`
  	Refresh_token *string

	Provider string	
}

type Role string
type SubscriptionPlan string

const (
	UserRole	 Role = "user"
	AdminRole	 Role = "admin"
	TeamLeadRole Role = "team_lead"
)

const (
	FreePlan SubscriptionPlan = "Free"
	ProPlan SubscriptionPlan = "Pro"
	EnterprisePlan SubscriptionPlan = "Enterprise"
)


func HashPassword(password string) (string, error){
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}


func VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
