package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID string

	CreatedAt time.Time
	UpdatedAt time.Time

	Email         string
	EmailVerified bool

	Account *Account

	Name   string
	Avatar *string

	Role             Role
	SubscriptionPlan SubscriptionPlan
}

type Account struct {
	UserID string
	PasswordHash  string
	Provider string
}

type Role string
type SubscriptionPlan string

const (
	UserRole     Role = "user"
	AdminRole    Role = "admin"
	TeamLeadRole Role = "team_lead"
)

const (
	FreePlan       SubscriptionPlan = "Free"
	ProPlan        SubscriptionPlan = "Pro"
	EnterprisePlan SubscriptionPlan = "Enterprise"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
