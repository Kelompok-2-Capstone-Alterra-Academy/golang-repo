package handler

import (
	"fmt"
	"net/http"

	"capston-lms/internal/application/service"
	"capston-lms/internal/application/usecase"
	"capston-lms/internal/entity"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	TransactionUsecase        usecase.TransactionUsecase
	Usecase                   usecase.UserUseCase
	TrasanctionDetailsUseCase usecase.TrasanctionDetailsUseCase
}

func (handler TransactionHandler) CheckoutTransaction() echo.HandlerFunc {
	return func(e echo.Context) error {
		var order entity.Transaction
		var orderDetails entity.TransactionDetails

		// Get UserID
		user := e.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwt.MapClaims)
		UserID := int((*claims)["id"].(float64))

		order.Status = "pending"
		order.StudentId = UserID

		err := handler.TransactionUsecase.CreateTransaction(order)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create user"})
		}

		if err := e.Bind(&orderDetails); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		// Set order_id for each order item and save to database
		orderID, err := handler.TransactionUsecase.GetLastTransactionID()

		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		orderDetails.TransactionId = orderID
		err = handler.TrasanctionDetailsUseCase.CreateOrderItems(orderDetails)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}

		// Create request body for Midtrans Snap API
		snapReq, err := handler.TransactionUsecase.GenerateSnapReq(orderID, UserID, order.TotalPayment)
		if err != nil {
				return e.JSON(500, echo.Map{
						"error": err.Error(),
				})
		}

		fmt.Println("================ Request with global config ================")
		service.SetupGlobalMidtransConfig()
		service.CreateTransactionWithGlobalConfig()

		fmt.Println("================ Request with Snap Client ================")
		service.InitializeSnapClient()
		respPayment, err := service.CreateTransaction(*snapReq)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}
		token := respPayment.Token
		redirectURL := respPayment.RedirectURL

		// Return response
		response := map[string]interface{}{
			"transaction":  snapReq,
			"tokenPayment": token,
			"redirectURL":  redirectURL,
		}
		return e.JSON(http.StatusOK, response)
	}
}
