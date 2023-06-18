package entity

import (
	"gorm.io/gorm"
)

type RateCourse struct {
	*gorm.Model

	
	UserId      *int    `json:"user_id" form:"user_id"`
	User        User  	`json:"user,omitempty" gorm:"foreignKey:UserId"`
	CourseId    *int    `json:"course_id" form:"course_id"`
	Course      Course  `json:"course,omitempty" gorm:"foreignKey:courseId"`
	Rating 		int 	`json:"rating" form:"rating"`
	Comment 	int 	`json:"comment" form:"comment"`
}
