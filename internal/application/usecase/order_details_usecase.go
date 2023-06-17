package usecase

import (
	"capston-lms/internal/adapters/repository"
	"capston-lms/internal/entity"
)

type TrasanctionDetailsUseCase struct {
	TransactionDetailRepo repository.TrasanctionDetailsRepository
}

func (usecase TrasanctionDetailsUseCase) CreateOrderItems(order entity.TransactionDetails) error {
	err := usecase.TransactionDetailRepo.CreateOrderItem(order)
	return err
}
func (usecase TrasanctionDetailsUseCase) GetOrderItemByBook(id int) ([]entity.TransactionDetails, error) {
	orderItems, err := usecase.TransactionDetailRepo.GetOrderItemsByBook(id)
	if err != nil {
		return nil, err
	}
	return orderItems, nil
}
