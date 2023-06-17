package repository

import (
	"capston-lms/internal/entity"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	DB *gorm.DB
}

func (repo *TransactionRepository) GetLastTransactionID() (uint, error) {
	fmt.Println("repo")

	var transaction entity.Transaction
	if err := repo.DB.Order("id desc").First(&transaction).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 1, nil
		}
		return 0, err
	}

	return transaction.ID, nil
}
func (r *TransactionRepository) FindByID(id uint) (*entity.Transaction, error) {
	var Transaction entity.Transaction
	err := r.DB.First(&Transaction, id).Error
	if err != nil {
		return nil, err
	}
	return &Transaction, nil
}

func (r *TransactionRepository) GetUserByID(id uint) (*entity.User, error) {
	var user entity.User
	err := r.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *TransactionRepository) GetCourseTransactionsByTransactionID(TransactionID uint) ([]entity.TransactionDetails, error) {
	var TrasanctionDetails []entity.TransactionDetails
	err := r.DB.Where("transaction_id = ?", TransactionID).Preload("Course").Find(&TrasanctionDetails).Error
	if err != nil {
		return nil, err
	}
	return TrasanctionDetails, nil
}
func (repo TransactionRepository) CreateTransaction(Transaction entity.Transaction) error {
	result := repo.DB.Create(&Transaction)
	return result.Error
}

func (repo TransactionRepository) UpdateTransaction(id int, Transaction entity.Transaction) error {
	result := repo.DB.Model(&Transaction).Where("id = ?", id).Updates(&Transaction)
	return result.Error
}