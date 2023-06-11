package entity

import (
	"gorm.io/gorm"
)

type Section struct {
	*gorm.Model
	CourseId   int    `json:"course_id" form:"course_id"`
	Course     Course `gorm:"foreignKey:CourseId"`
	CourseName string `json:"course_name" form:"course_name"`
}
