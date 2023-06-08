package entity

import "gorm.io/gorm"

type TaskMaterial struct {
	*gorm.Model

	TaskName         string `json:"task_name" form:"task_name" validate:"required"`
	Description      string `json:"description" form:"description" validate:"required"`
	Type             string `json:"type" form:"type" validate:"required"`
	AttachmentSource string `json:"attachment_source" form:"attachment_source" validate:"required"`
	Status           string `json:"status" gorm:"default:students"`
}
