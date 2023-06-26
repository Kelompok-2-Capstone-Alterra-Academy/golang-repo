package entity

import (
	"gorm.io/gorm"
)

type EducationNews struct {
	*gorm.Model

	UserId           uint   `json:"user_id" form:"user_id"`
	Tittle           string `json:"tittle" form:"tittle" validate:"required"`
	ShortDescription string `json:"short_description" form:"short_description" validate:"required" `
	Thumbnail        string `json:"thumbnail" form:"thumbnail" validate:"required"`
	Link             string `json:"link" form:"Link" validate:"required"`
}
