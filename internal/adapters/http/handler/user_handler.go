package handler

import (
	"net/http"
	"strconv"

	"capston-lms/internal/application/service"
	"capston-lms/internal/application/usecase"
	"capston-lms/internal/entity"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	UserUsecase usecase.UserUseCase
}

func (handler UserHandler) GetAllUsers() echo.HandlerFunc {
	return func(e echo.Context) error {
		var users []entity.User

		users, err := handler.UserUsecase.GetAllUsers()
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "success get all user",
			"data":        users,
		})
	}
}

func (handler UserHandler) GetUser() echo.HandlerFunc {
	return func(e echo.Context) error {
		var user entity.User
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     err.Error(),
			})
		}

		user, err = handler.UserUsecase.GetUser(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "success get user by id",
			"data":        user,
		})
	}
}

func (handler UserHandler) CreateUser() echo.HandlerFunc {
	return func(e echo.Context) error {
		var user entity.User
		if err := e.Bind(&user); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     err.Error(),
			})
		}

		// Validasi input menggunakan package validator
		validate := validator.New()
		if err := validate.Struct(user); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     err.Error(),
			})
		}

		// Validasi email unik
		if err := handler.UserUsecase.UniqueEmail(user.Email); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     err.Error(),
			})
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     "failed to created user",
			})
		}
		user.Password = string(hashedPassword)

		err = handler.UserUsecase.CreateUser(user)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     "failed to created user",
			})
		}

		return e.JSON(http.StatusCreated, map[string]interface{}{
			"status code": http.StatusCreated,
			"message":     "success create new user",
			"data":        user,
		})
	}
}
func (handler UserHandler) GetUserByRole() echo.HandlerFunc {
	return func(e echo.Context) error {
		var user []entity.User
		role := "mentors"
		user, err := handler.UserUsecase.GetUserByRole(role)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "success get user by role",
			"data":        user,
		})
	}
}

func (handler UserHandler) DeleteUser() echo.HandlerFunc {
	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     "input is not a number",
			})
		}

		err = handler.UserUsecase.DeleteUser(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     "failed to created user",
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "success delete data",
		})
	}
}

func (handler UserHandler) UpdateUser() echo.HandlerFunc {
	var user entity.User

	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     err.Error(),
			})
		}

		// Check if profile image is uploaded
		profileFile, err := e.FormFile("profile")
		if err == nil {
			// If uploaded, process the file
			src, err := profileFile.Open()
			if err != nil {
				return err
			}
			defer src.Close()

			// Upload the file to S3 or any other storage service
			uploadResult, err := service.UploadToS3(e, profileFile.Filename, src)
			if err != nil {
				return err
			}

			// Update the user's profile image
			user.Profile = uploadResult
		}
		err = handler.UserUsecase.UpdateUser(id, user)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		if err := e.Bind(&user); err != nil {
			return e.JSON(http.StatusNotFound, map[string]interface{}{
				"status code": http.StatusNotFound,
				"message":     err.Error(),
			})
		}

		err = handler.UserUsecase.UpdateUser(id, user)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "success update user data",
			"data":        user,
		})
	}
}
