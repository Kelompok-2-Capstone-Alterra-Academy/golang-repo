package entity

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	*gorm.Model
	DueDate  time.Time `json:"due_date"`
	ModuleId string    `json:"module_id" form:"module_id"`
	Module   Module    `json:"module" gorm:"foreignKey:ModuleId"`
}
