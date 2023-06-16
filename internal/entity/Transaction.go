package entity

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	*gorm.Model

	InvoiceNumber string `json:"invoice_number" form:"invoice_number"`
	TotalPayment  int    `json:"total_payment" form:"total_payment"`
	AdminFees     int    `json:"admin_fees" form:"admin_fees"`
	Status        string `json:"status" form:"status"`
	StudentId     int    `json:"student_id" form:"student_id"`
	Student       User   `json:"student,omitempty" gorm:"foreignKey:StudentId"`
	PromoId       *int   `json:"promo_id" form:"promo_id"`
	Promo         Promo  `json:"promo,omitempty" gorm:"foreignKey:PromoId"`
}

func (t *Transaction) GenerateInvoiceNumber() {
	currentTime := time.Now().Format("20060102")
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(99999)

	t.InvoiceNumber = fmt.Sprintf("%s-%04d", currentTime, randomNumber)
}
