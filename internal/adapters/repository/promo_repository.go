package repository

import (
	"capston-lms/internal/entity"

	"gorm.io/gorm"
)

type PromoRepository struct {
	DB *gorm.DB
}

func (repo PromoRepository) GetAllPromo() ([]entity.Promo, error) {
	var promos []entity.Promo
	result := repo.DB.Find(&promos)
	return promos, result.Error
}

func (repo PromoRepository) GetPromo(id int) (entity.Promo, error) {
	var promos entity.Promo
	result := repo.DB.First(&promos, id)
	return promos, result.Error
}

func (repo PromoRepository) CreatePromo(promo entity.Promo) error {
	result := repo.DB.Create(&promo)
	return result.Error
}

func (repo PromoRepository) UpdatePromo(id int, promo entity.Promo) error {
	result := repo.DB.Model(&promo).Where("id = ?", id).Updates(&promo)
	return result.Error
}

func (repo PromoRepository) DeletePromo(id int) error {
	result := repo.DB.Delete(&entity.Promo{}, id)
	return result.Error
}

func (repo PromoRepository) FindPromo(id int) error {
	result := repo.DB.First(&entity.Promo{}, id)
	return result.Error
}