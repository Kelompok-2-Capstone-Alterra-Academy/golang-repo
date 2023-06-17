package entity

import (
	"gorm.io/gorm"
)

type Transaction struct {
	*gorm.Model

	InvoiceNumber      string               `json:"invoice_number" form:"invoice_number"`
	TotalPayment       int                  `json:"total_payment" form:"total_payment"`
	AdminFees          int                  `json:"admin_fees" form:"admin_fees"`
	Status             string               `json:"status" form:"status"`
	StudentId          int                  `json:"student_id" form:"student_id"`
	Student            User                 `json:"student,omitempty" gorm:"foreignKey:StudentId"`
	PromoId            *int                 `json:"promo_id" form:"promo_id"`
	Promo              Promo                `json:"promo,omitempty" gorm:"foreignKey:PromoId"`
	Token              string               `json:"token" form:"token"`
	TransactionDetails []TransactionDetails `json:"transaction_details,omitempty" gorm:"foreignKey:TransactionId"`
}
