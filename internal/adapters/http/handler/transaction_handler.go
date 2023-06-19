package handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"capston-lms/internal/application/service"
	"capston-lms/internal/application/usecase"
	"capston-lms/internal/entity"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type TransactionHandler struct {
	TransactionUsecase        usecase.TransactionUsecase
	Usecase                   usecase.UserUseCase
	TrasanctionDetailsUseCase usecase.TrasanctionDetailsUseCase
}

func (handler TransactionHandler) GetMyTransaction() echo.HandlerFunc {
	return func(e echo.Context) error {
		var Transaction []entity.Transaction
		StudentId, err := service.GetUserIDFromToken(e)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		Transaction, err = handler.TransactionUsecase.GetTransaction(StudentId)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}
		data := make(map[string]interface{})
		data["transaction"] = Transaction
		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "success get Transaction history ",
			"data":        data,
		})
	}
}
func (handler TransactionHandler) MidtransNotification() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Ambil data notifikasi dari body request
		var jsonReq struct {
			OrderID string `json:"order_id"`
			Status  string `json:"transaction_status"`
		}
		if err := c.Bind(&jsonReq); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}
		invoice_number := jsonReq.OrderID

		var Transaction entity.Transaction
		var CourseEnrollment entity.CourseEnrollment

		Transaction, err := handler.TransactionUsecase.FindByInvoiceId(invoice_number)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		// Lakukan pemrosesan notifikasi sesuai kebutuhan Anda
		// Anda dapat menggunakan package encoding/json untuk mendecode JSON payload

		// Contoh sederhana: Cetak payload notifikasi
		fmt.Println("Received Midtrans notification:")

		// Lakukan verifikasi signature notifikasi dari Midtrans
		// Implementasikan logika verifikasi sesuai dokumentasi Midtrans

		// Contoh sederhana: Verifikasi selalu berhasil
		validSignature := true

		if validSignature {
			// Verifikasi signature berhasil

			// Periksa status transaksi
			transactionStatus := jsonReq.Status

			switch transactionStatus {
			case "settlement":
				// Transaksi berhasil
				CourseEnrollment.UserId = Transaction.StudentId
				CourseEnrollment.CourseId = Transaction.TransactionDetails[0].CourseId
				err := handler.TransactionUsecase.CreateEnrolment(CourseEnrollment)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create user"})
				}
				Transaction.Status = "success"
				err = handler.TransactionUsecase.UpdateTransaction(int(Transaction.ID), Transaction)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, map[string]string{
						"error": err.Error(),
					})
				}
				return c.JSON(http.StatusOK, "Transaction successfully captured")

			case "expire":
				// Transaksi kedaluwarsa

				// Lakukan penanganan transaksi yang kedaluwarsa pada website Anda
				Transaction.Status = "expire"
				err = handler.TransactionUsecase.UpdateTransaction(int(Transaction.ID), Transaction)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, map[string]string{
						"error": err.Error(),
					})
				}
				return c.JSON(http.StatusOK, "Transaction expired")

			default:
				// Status transaksi lainnya

				// Tangani status transaksi lain sesuai kebutuhan Anda
				Transaction.Status = "pending"
				err = handler.TransactionUsecase.UpdateTransaction(int(Transaction.ID), Transaction)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, map[string]string{
						"error": err.Error(),
					})
				}
				return c.JSON(http.StatusOK, "Other transaction status")
			}
		} else {
			// Verifikasi signature gagal
			log.Error("Invalid signature")

			return c.JSON(http.StatusBadRequest, "Invalid signature")
		}
	}
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
		currentTime := time.Now().Format("20060102")
		rand.Seed(time.Now().UnixNano())
		randomNumber := rand.Intn(99999)
		order.InvoiceNumber = fmt.Sprintf("%s%04d", currentTime, randomNumber)

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

		order.Token = respPayment.RedirectURL
		err = handler.TransactionUsecase.UpdateTransaction(int(orderID), order)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}
		// Return response
		response := map[string]interface{}{
			"transaction":  snapReq,
			"tokenPayment": token,
			"redirectURL":  redirectURL,
		}
		return e.JSON(http.StatusOK, response)
	}
}
