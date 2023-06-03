package entity

import (
	"gorm.io/gorm"
)

type CourseSection struct {
	*gorm.Model

	CourseId   string `json:"course_id" form:"course_id"`
	Course     Course `gorm:"foreignKey:CourseId"`
	CourseName string `json:"course_name" form:"course_name"`
}
