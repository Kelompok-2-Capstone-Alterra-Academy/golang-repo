package repository

import (
	"capston-lms/internal/entity"

	"gorm.io/gorm"
)

type TrasanctionDetailsRepository struct {
	DB *gorm.DB
}

func (repo TrasanctionDetailsRepository) CreateOrderItem(order entity.TransactionDetails) error {
	result := repo.DB.Create(&order)
	return result.Error
}

func (repo TrasanctionDetailsRepository) GetOrderItemsByBook(id int) ([]entity.TransactionDetails, error) {
	var orderItems []entity.TransactionDetails
	result := repo.DB.Where("course_id = ?", id).Find(&orderItems)
	return orderItems, result.Error
}
