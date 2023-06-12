package entity

import "gorm.io/gorm"

type Section struct {
	*gorm.Model
	SectionName string `json:"section_name" form:"section_name"`
	CourseId    string `json:"course_id" form:"course_id"`
	Course      Course `json:"course" gorm:"foreignKey:CourseId"`
}