package entity

import (
	"gorm.io/gorm"
)

type Task struct {
	*gorm.Model
	DueDate  string  `json:"due_date" form:"due_date"`
	ModuleId int     `json:"module_id" form:"module_id"`
	Module   *Module `json:"module,omitempty" gorm:"foreignKey:ModuleId"`
}
