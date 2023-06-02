package entity

import (
	"gorm.io/gorm"
)

type Major struct {
	*gorm.Model

	MajorName string `json:"major_name" form:"major_name"`
}
