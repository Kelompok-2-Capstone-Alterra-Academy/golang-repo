package entity

import (
	"mime/multipart"

	"gorm.io/gorm"
)

type TaskSubmission struct {
	*gorm.Model
	Notes string                `json:"notes" form:"notes"`
	File  *multipart.FileHeader `json:"file" form:"file"`
}
