package usecase

import (
	"capston-lms/internal/adapters/repository"
	"capston-lms/internal/entity"
)

type PromoUseCase struct {
	Repo repository.PromoRepository
}

func (usecase PromoUseCase) GetAllPromo() ([]entity.Promo, error) {
	promos, err := usecase.Repo.GetAllPromo()
	return promos, err
}

func (usecase PromoUseCase) GetPromo(id int) (entity.Promo, error) {
	promo, err := usecase.Repo.GetPromo(id)
	return promo, err
}

func (usecase PromoUseCase) CreatePromo(promo entity.Promo) error {
	err := usecase.Repo.CreatePromo(promo)
	return err
}

func (usecase PromoUseCase) UpdatePromo(id int, promo entity.Promo) error {
	err := usecase.Repo.UpdatePromo(id, promo)
	return err
}

func (usecase PromoUseCase) DeletePromo(id int) error {
	err := usecase.Repo.DeletePromo(id)
	return err
}

func (usecase PromoUseCase) FindPromo(id int) error {
	err := usecase.Repo.FindPromo(id)
	return err
}