package entity

import "gorm.io/gorm"

type Task struct {
	*gorm.Model
	Title     string  `json:"title" form:"title"`
	DueDate   string  `json:"due_date" form:"due_date"`
	SectionId string  `json:"section_id" form:"section_id"`
	Course    Section `json:"section" gorm:"foreignKey:SectionId"`
}
