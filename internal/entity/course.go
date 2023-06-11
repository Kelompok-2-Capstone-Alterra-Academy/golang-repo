package entity

import (
	"gorm.io/gorm"
)

type Course struct {
	*gorm.Model

	CategoryId      int      `json:"category_id" form:"category_id" validate:"required"`
	Category        Category `gorm:"foreignKey:CategoryId"`
	ClassId         int      `json:"class_id" form:"class_id" validate:"required"`
	Class           Class    `gorm:"foreignKey:ClassId"`
	MentorId        int      `json:"mentor_id" form:"mentor_id" validate:"required"`
	Mentor          User     `gorm:"foreignKey:MentorId"`
	MajorId         int      `json:"major_id" form:"major_id" validate:"required"`
	Major           Major    `gorm:"foreignKey:MajorId"`
	CourseName      string   `json:"course_name" form:"course_name" validate:"required"`
	Price           string   `json:"price" form:"price" validate:"required"`
	Duration        string   `json:"duration" form:"duration" validate:"required"`
	Status          string   `json:"status" form:"status" validate:"required"`
	Description     string   `json:"description" form:"description" validate:"required"`
	Thumbnail       string   `json:"thumbnail" form:"thumbnail" validate:"required"`
	LiveSessionWeek string   `json:"live_session_week" form:"live_session_week" validate:"required"`
}
