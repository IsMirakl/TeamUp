package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID uint	`gorm:"primaryKey"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Email string `gorm:"unique;not null"`
	EmailVerified bool

	Account *Account `gorm:"foreignKey:UserID"`

	Name string `gorm:"size:25;not null;"`
	Avatar *string

	Role Role	`gorm:"type:varchar(20);default:'user';check:role IN ('user', 'admin', 'team_lead')"`
	SubscriptionPlan SubscriptionPlan `gorm:"type:varchar(20);default:'Free';check:subscription_plan IN ('Free', 'Pro', 'Enterprise')"`
}

type Account struct {
	ID string `gorm:"primaryKey"`

	UserID uint `gorm:"index"`
	User User

	passwordHash  string `gorm:"size255"`
  	refresh_token *string
  	access_token  *string

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


func HashPassword(password string) ([]byte, error){
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}


func VerifyPassword(hashedPassword, password string) (error){
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}