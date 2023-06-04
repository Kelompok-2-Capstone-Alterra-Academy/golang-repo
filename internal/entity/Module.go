package entity

import "gorm.io/gorm"

type Module struct {
	*gorm.Model

	Name        string `json:"name" form:"name" validate:"required"`
	Email       string `json:"email" form:"email" validate:"required,email" `
	Password    string `json:"password" form:"password" validate:"required"`
	Role        string `gorm:"default:students"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required"`
}
