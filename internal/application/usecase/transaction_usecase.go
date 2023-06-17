package usecase

import (
	"capston-lms/internal/adapters/repository"
	"capston-lms/internal/entity"
	"fmt"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type TransactionUsecase struct {
	TransactionRepo repository.TransactionRepository
	UserRepo        repository.UserRepository
}

func (usecase TransactionUsecase) GetTransaction(id int) ([]entity.Transaction, error) {
	transaction, err := usecase.TransactionRepo.GetTransaction(id)
	return transaction, err
}

func (uc *TransactionUsecase) GetLastTransactionID() (uint, error) {
	stores, err := uc.TransactionRepo.GetLastTransactionID()
	return stores, err
}

func (usecase TransactionUsecase) CreateTransaction(user entity.Transaction) error {
	err := usecase.TransactionRepo.CreateTransaction(user)
	return err
}

func (usecase TransactionUsecase) UpdateTransaction(id int, Transaction entity.Transaction) error {
	err := usecase.TransactionRepo.UpdateTransaction(id, Transaction)
	return err
}

func (uc *TransactionUsecase) GenerateSnapReq(TransactionID uint, UserID int, TotalPrice int) (*snap.Request, error) {
	// Get the order and its related data from the repository
	order, err := uc.TransactionRepo.FindByID(TransactionID)
	if err != nil {
		return nil, err
	}
	user, err := uc.UserRepo.GetUser(UserID)
	if err != nil {
		return nil, err
	}

	courseOrders, err := uc.TransactionRepo.GetCourseTransactionsByTransactionID(order.ID)
	if err != nil {
		return nil, err
	}

	// Set customer detail data
	custAddress := &midtrans.CustomerAddress{
		FName:       user.Name,
		LName:       "Doe",
		Phone:       "083848988030",
		Address:     "Malang dinoyo",
		City:        "Malang",
		Postcode:    "16000",
		CountryCode: "IDN",
	}
	custDetail := &midtrans.CustomerDetails{
		FName:    user.Name,
		LName:    "- ",
		Email:    user.Email,
		Phone:    "083848988030",
		BillAddr: custAddress,
		ShipAddr: custAddress,
	}

	// Create ItemDetails array for Snap Request
	var itemDetails []midtrans.ItemDetails
	var totalPrice int64 = 0
	for _, bo := range courseOrders {
		itemDetails = append(itemDetails, midtrans.ItemDetails{
			ID:    bo.CourseId,
			Price: int64(bo.Price),
			Qty:   int32(1),
			Name:  bo.Course.CourseName,
		})
		totalPrice += int64(bo.Price) * int64(1)
	}

	// Create Snap Request object
	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  fmt.Sprint(TransactionID),
			GrossAmt: int64(totalPrice),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail:  custDetail,
		EnabledPayments: snap.AllSnapPaymentType,
		Items:           &itemDetails,
	}

	return snapReq, nil
}
