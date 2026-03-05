package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID uint	`gorm:"primaryKey"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Email string `gorm:"unique;not null"`
	EmailVerified bool


	Name string `gorm:"size:25;not null;"`
	Avatar *string

	Role Role	`gorm:"type:varchar(20);default:'user'; check:role IN ('user', 'admin', 'team_lead')"`
	SubscriptionPlan SubscriptionPlan `gorm:"type:varchar(20);default:'Free'; check:subscription_plan IN ('Free', 'Pro', 'Enterprise')"`
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