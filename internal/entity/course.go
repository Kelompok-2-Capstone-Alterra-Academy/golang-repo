package entity

import (
	"gorm.io/gorm"
)

type Course struct {
	*gorm.Model

	CategoryId      int   `json:"category_id" form:"category_id"`
	Category        Category `gorm:"foreignKey:CategoryId"`
	ClassId         int   `json:"class_id" form:"class_id"`
	Class           Class    `gorm:"foreignKey:ClassId"`
	MentorId        int   `json:"mentor_id" form:"mentor_id"`
	Mentor          User     `gorm:"foreignKey:MentorId"`
	MajorId         int   `json:"major_id" form:"major_id"`
	Major           Major    `gorm:"foreignKey:MajorId"`
	CourseName      string   `json:"course_name" form:"course_name"`
	Price           string   `json:"price" form:"price"`
	Duration        string   `json:"duration" form:"duration"`
	Status          string   `json:"status" form:"status"`
	Description     string   `json:"description" form:"description"`
	Thumbnail       string   `json:"thumbnail" form:"thumbnail"`
	LiveSessionWeek string   `json:"live_session_week" form:"live_session_week"`
