package entity

import (
	"gorm.io/gorm"
)

type CourseEnrollment struct {
	*gorm.Model

	CourseId string `json:"course_id" form:"course_id"`
	UserId   string `json:"user_id" form:"user_id"`
	Course   Course `json:"course" gorm:"foreignKey:CourseId"`
	User     User   `json:"user" gorm:"foreignKey:UserId"`
	Status   string `json:"status" form:"status"`
}
