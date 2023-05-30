package entity

import (
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model

	Name        string `json:"name" form:"name" validate:"required"`
	Email       string `json:"email" form:"email" validate:"required,email" `
	Password    string `json:"password" form:"password" validate:"required"`
	Role        string `gorm:"default:students"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Status      string `gorm:"default:not-verified"`
	SchoolName  string `json:"school_name" form:"school_name"`
	Class       string `json:"class" form:"class"`
	Gender      string `json:"gender" form:"gender"`
}
