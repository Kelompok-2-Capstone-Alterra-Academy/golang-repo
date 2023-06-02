package entity

import (
	"time"

	"gorm.io/gorm"
)

type OTPToken struct {
	gorm.Model
	Token     string
	Email     string
	ExpiredAt time.Time
}
