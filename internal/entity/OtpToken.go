package entity

import (
	"time"

	"gorm.io/gorm"
)

type OTPToken struct {
	gorm.Model
	Otp       string `json:"otp" form:"otp"`
	Email     string `json:"email" form:"email"`
	ExpiredAt time.Time
}
