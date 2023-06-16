package entity

import (
	"gorm.io/gorm"
)

type Modules struct {
	*gorm.Model

	CourseSectionId string        `json:"course_section_id" form:"course_section_id"`
	CourseSection   CourseSection `gorm:"foreignKey:CourseSectionId"`
	Description     string        `json:"description" form:"description" validate:"required"`
	ModuleName      string        `json:"module_name" form:"module_name" validate:"required"`
	Type            string        `json:"type" form:"type" validate:"required"`
	Link            string        `json:"link" form:"link" validate:"required"`
}
