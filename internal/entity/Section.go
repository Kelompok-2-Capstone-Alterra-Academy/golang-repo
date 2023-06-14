package entity

import "gorm.io/gorm"

type Section struct {
	*gorm.Model
	SectionName string `json:"section_name" form:"section_name"`
	CourseId    int    `json:"course_id" form:"course_id"`
}
