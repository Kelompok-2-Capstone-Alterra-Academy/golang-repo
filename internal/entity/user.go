package entity

import (
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model

	Name        string `json:"name" form:"name" `
	Email       string `json:"email" form:"email" `
	Password    string `json:"password" form:"password" `
	Role        string `json:"role" gorm:"default:students"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Status      string `json:"status" gorm:"default:not-verified"`
	SchoolName  string `json:"school_name" form:"school_name"`
	Class       string `json:"class" form:"class"`
	Gender      string `json:"gender" form:"gender"`
	Profile     string `json:"profile" form:"profile"`
}
