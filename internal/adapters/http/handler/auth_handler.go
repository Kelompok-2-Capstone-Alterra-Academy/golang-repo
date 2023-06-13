package handler

import (
	"encoding/json"
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
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status_code": http.StatusBadRequest,
				"message":     "Invalid request body",
			})
		}

		// Validasi input menggunakan package validator
		validate := validator.New()
		if err := validate.Struct(user); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status_code": http.StatusBadRequest,
				"message":     "Validation errors",
				"errors":      err.Error(),
			})
		}

		// Validasi email unik
		if err := handler.Usecase.UniqueEmail(user.Email); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status_code": http.StatusBadRequest,
				"message":     "Validation errors",
				"errors":      err.Error(),
			})
		}

		hashedPassword, err := service.Encrypt(user.Password)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create user"})
		}
		user.Password = string(hashedPassword)
		user.Role = "students"
		user.Status = "not-verified"

		// sending otp
		otp := service.GenerateOTP()
		// Simpan token ke database
		expiredAt := time.Now().Add(time.Minute * 5) // Token berlaku selama 5 menit
		otpToken := entity.OTPToken{
			Otp:       otp,
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
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status_code": http.StatusInternalServerError,
				"message":     "Failed to send OTP email",
				"errors":      err.Error(),
			})
		}

		err = handler.Usecase.CreateUser(user)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create user"})
		}

		data := make(map[string]interface{})
		data["users"] = user
		return e.JSON(http.StatusCreated, map[string]interface{}{
			"status_code": http.StatusCreated,
			"message":     "user created successfully",
			"data":        data,
		})
	}
}

func (handler AuthHandler) NewPassword() echo.HandlerFunc {
	return func(e echo.Context) error {
		var user entity.User

		if err := e.Bind(&user); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status_code": http.StatusBadRequest,
				"message":     "Invalid request body",
			})
		}

		hashedPassword, err := service.Encrypt(user.Password)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create user"})
		}
		user.Password = string(hashedPassword)
		StudentId, err := service.GetUserIDFromToken(e)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}
		err = handler.Usecase.UpdateUser(StudentId, user)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create user"})
		}

		data := make(map[string]interface{})
		data["users"] = user
		return e.JSON(http.StatusCreated, map[string]interface{}{
			"status_code": http.StatusCreated,
			"message":     "user created successfully",
			"data":        data,
		})
	}
}

func (handler AuthHandler) ForgotPassword() echo.HandlerFunc {
	return func(e echo.Context) error {
		var user entity.User
		if err := e.Bind(&user); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status_code": http.StatusBadRequest,
				"message":     "Invalid request body",
			})
		}

		// sending otp
		otp := service.GenerateOTP()
		// Simpan token ke database
		expiredAt := time.Now().Add(time.Minute * 5) // Token berlaku selama 5 menit
		otpToken := entity.OTPToken{
			Otp:       otp,
			Email:     user.Email,
			ExpiredAt: expiredAt,
		}
		err := handler.Usecase.SaveOTP(otpToken)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to save otp token"})
		}

		body := "OTP Kamu adalah sebagai berikut ini : " + otp
		err = service.SendEmail(user.Email, "lakukan verifikasi akun anda sebelum 10 menit", body)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status_code": http.StatusInternalServerError,
				"message":     "Failed to send OTP email",
				"errors":      err.Error(),
			})
		}

		data := make(map[string]interface{})
		data["users"] = user
		return e.JSON(http.StatusCreated, map[string]interface{}{
			"status_code": http.StatusCreated,
			"message":     "user created successfully",
			"data":        data,
		})
	}
}
func (handler AuthHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Bind request body to user struct
		var user entity.User
		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"status_code": http.StatusBadRequest,
				"message":     "Invalid request body",
			})
		}

		// Get user by email
		dbUser, err := handler.Usecase.GetUserByEmail(user.Email)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"status_code": http.StatusUnauthorized,
				"message":     "Invalid email or password",
			})
		}

		// Check password
		if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {

			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"status_code": http.StatusUnauthorized,
				"message":     "Invalid email or password",
			})
		}
		// Check if user status is "not-verified"
		if dbUser.Status == "not-verified" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"status_code": http.StatusUnauthorized,
				"message":     "User is not verified",
			})
		}
		t, err := middleware.CreateToken(int(dbUser.ID), dbUser.Email, dbUser.Role)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status_code": http.StatusInternalServerError,
				"message":     "Failed to create token",
			})
		}

		data := make(map[string]interface{})
		data["token"] = t
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status_code": http.StatusOK,
			"message":     "congratulations successful login",
			"data":        data,
		})
	}
}

func (handler AuthHandler) VerifyOTP() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Read the request body
		body := c.Request().Body

		// Create an instance of OTPToken struct
		var otpToken entity.OTPToken

		// Decode the JSON request body into otpToken struct
		if err := json.NewDecoder(body).Decode(&otpToken); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"status_code": http.StatusBadRequest,
				"message":     "Invalid request body",
			})
		}

		email := otpToken.Email
		otp := otpToken.Otp
		result := handler.Usecase.VerifiedOtpToken(email, otp)
		if result != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"status_code": http.StatusBadRequest,
				"message":     "Invalid OTP token",
			})
		}

		dbUser, err := handler.Usecase.GetUserByEmail(email)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"status_code": http.StatusUnauthorized,
				"message":     "Invalid email or password",
			})
		}

		t, err := middleware.CreateToken(int(dbUser.ID), dbUser.Email, dbUser.Role)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status_code": http.StatusInternalServerError,
				"message":     "Failed to generate token",
			})
		}

		data := make(map[string]interface{})
		data["token"] = t

		return c.JSON(http.StatusOK, map[string]interface{}{
			"status_code": http.StatusOK,
			"message":     "OTP token has been verified",
			"data":        data,
		})
	}
}
