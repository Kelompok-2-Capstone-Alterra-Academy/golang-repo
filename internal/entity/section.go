package entity

import (
	"gorm.io/gorm"
)

type Section struct {
	*gorm.Model
	CourseId 		int `json:"course_id" form:"course_id"`
	Course 			Course `gorm:"foreignKey:CourseId"`
	SectionName  	string `json:"section_name" form:"section_name"`
}
