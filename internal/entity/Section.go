package entity

import "gorm.io/gorm"

type Section struct {
	*gorm.Model
	CourseName string `json:"course_name" form:"course_name"`
	CourseId   string `json:"course_id" form:"course_id"`
	Course     Course `json:"course" gorm:"foreignKey:CourseId"`
}
