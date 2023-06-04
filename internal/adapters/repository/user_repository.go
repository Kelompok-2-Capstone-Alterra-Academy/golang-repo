package repository

import (
	"capston-lms/internal/entity"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (repo UserRepository) GetAllUsers() ([]entity.User, error) {
	var users []entity.User
	result := repo.DB.Find(&users)
	return users, result.Error
}

func (repo UserRepository) GetUser(id int) (entity.User, error) {
	var users entity.User
	result := repo.DB.First(&users, id)
	return users, result.Error
}

func (repo UserRepository) CreateUser(user entity.User) error {
	result := repo.DB.Create(&user)
	return result.Error
}

func (repo UserRepository) UpdateUser(id int, user entity.User) error {
	result := repo.DB.Model(&user).Where("id = ?", id).Updates(&user)
	return result.Error
}

func (repo UserRepository) DeleteUser(id int) error {
	result := repo.DB.Delete(&entity.User{}, id)
	return result.Error
}

func (repo UserRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := repo.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (repo UserRepository) UniqueEmail(email string) error {
	var user entity.User
	result := repo.DB.Where("email = ?", email).First(&user)
	if result.RowsAffected > 0 {
		return fmt.Errorf("email %s already exists", email)
	}
	return nil
}

func (repo UserRepository) SaveOTP(otp entity.OTPToken) error {
	result := repo.DB.Create(&otp)
	return result.Error
}

func (repo UserRepository) VerifiedOtpToken(email string, token string) error {
	var otpToken entity.OTPToken
	err := repo.DB.Where("email = ? AND otp = ? AND status = ? ", email, token, "not-used").First(&otpToken).Error
	if err != nil {
		return err
	}

	err = repo.DB.Model(&otpToken).Update("status", "used").Error
	if err != nil {
		return err
	}

	user := entity.User{}
	err = repo.DB.Model(&user).Where("email = ?", email).Update("status", "active").Error
	if err != nil {
		return err
	}

	return nil
}
