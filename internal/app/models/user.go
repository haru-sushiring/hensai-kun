package models

import "time"

type User struct {
	ID uint `gorm:"primaryKey"`
	Email string `gorm:"unique;not null"`
	EmailVerifiedAt time.Time
	Password string `gorm:"not null"`
}