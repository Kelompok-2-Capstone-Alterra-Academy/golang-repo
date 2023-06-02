package entity

import (
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model

	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email" `
	Password    string `json:"password" form:"password"`
	Role        string `gorm:"default:students"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
}
