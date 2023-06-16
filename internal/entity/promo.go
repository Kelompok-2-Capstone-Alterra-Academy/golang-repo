package entity

import (
	"time"

	"gorm.io/gorm"
)

type Promo struct {
	*gorm.Model

	PromoName   string    `json:"promo_name" form:"promo_name"`
	ExpiredDate time.Time `json:"expired_date" form:"expired_date"`
	TotalPromo  float64   `json:"total_promo" form:"total_promo"`
	Thumbnail   string    `json:"thumbnail" form:"thumbnail" validate:"required"`
}
