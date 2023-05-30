package handler

import (
	"net/http"
	"time"

	"capston-lms/internal/adapters/http/middleware"
	"capston-lms/internal/application/service"
	"capston-lms/internal/application/usecase"

	"capston-lms/internal/entity"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	Usecase usecase.UserUseCase
}

func (handler AuthHandler) Register() echo.HandlerFunc {
	return func(e echo.Context) error {
		var user entity.User
		if err := e.Bind(&user); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
		}

		// Validasi input menggunakan package validator
		validate := validator.New()
		if err := validate.Struct(user); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Validation errors", "errors": err.Error()})
		}

		// Validasi email unik
		if err := handler.Usecase.UniqueEmail(user.Email); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}

		hashedPassword, err := service.Encrypt(user.Password)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create user"})
		}
		user.Password = string(hashedPassword)
		user.Role = "students"

		// sending otp
		otp := service.GenerateOTP()
		// Simpan token ke database
		expiredAt := time.Now().Add(time.Minute * 5) // Token berlaku selama 5 menit
		otpToken := entity.OTPToken{
			Token:     otp,
			Email:     user.Email,
			ExpiredAt: expiredAt,
		}
		err = handler.Usecase.SaveOTP(otpToken)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to save otp token"})
		}

		body := "OTP Kamu adalah sebagai berikut ini : " + otp
		err = service.SendEmail(user.Email, "lakukan verifikasi akun anda sebelum 10 menit", body)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to send OTP email")
		}

		err = handler.Usecase.CreateUser(user)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create user"})
		}

		return e.JSON(http.StatusCreated, user)
	}
}

func (handler AuthHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Bind request body to user struct
		var user entity.User
		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
		}

		// Get user by email
		dbUser, err := handler.Usecase.GetUserByEmail(user.Email)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid email or password"})
		}

		// Check password
		if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid email or password"})
		}

		t, err := middleware.CreateToken(int(dbUser.ID), dbUser.Email, dbUser.Role)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create token"})
		}

		return c.JSON(http.StatusOK, map[string]string{"token": t})
	}
}
