package entity

import (
	"gorm.io/gorm"
)

type Class struct {
	*gorm.Model

	ClassName        string `json:"class_name" form:"class_name"`
}
