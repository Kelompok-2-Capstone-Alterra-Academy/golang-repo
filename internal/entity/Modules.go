package entity

import "gorm.io/gorm"

type Attachment struct {
	*gorm.Model

	AttachmentName   string `json:"attachment_name" form:"attachment_name" validate:"required"`
	Description      string `json:"description" form:"description" validate:"required"`
	Type             string `json:"type" form:"type" validate:"required"`
	AttachmentSource string `json:"attachment_source" form:"attachment_source" validate:"required"`
	Status           string `json:"status" gorm:"default:students"`
	MentorId         string `json:"mentor_id" form:"mentor_id"`
	Mentor           User   `json:"mentor" gorm:"foreignKey:MentorId"`
}
