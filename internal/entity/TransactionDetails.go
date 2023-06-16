package entity

import (
	"gorm.io/gorm"
)

type TransactionDetails struct {
	*gorm.Model

	Price         int         `json:"price" form:"price"`
	TransactionId uint        `json:"transaction_id" form:"transaction_id"`
	Transaction   Transaction `json:"transaction" gorm:"foreignKey:TransactionId"`
	CourseId      string      `json:"course_id" form:"course_id"`
	Course        *Course     `json:"course,omitempty" gorm:"foreignKey:CourseId"`
}
