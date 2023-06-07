package handler

import (
	"net/http"
	"strconv"

	"capston-lms/internal/application/usecase"
	"capston-lms/internal/entity"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type MajorHandler struct {
	MajorUsecase usecase.MajorUseCase
}

func (handler MajorHandler) GetAllMajors() echo.HandlerFunc {
	return func(e echo.Context) error {
		var majors []entity.Major

		majors, err := handler.MajorUsecase.GetAllMajors()
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message": "success get all major",
			"data":   majors,
		})
	}
}

func (handler MajorHandler) GetMajor() echo.HandlerFunc {
	return func(e echo.Context) error {
		var major entity.Major
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message": err.Error(),
			})
		}

		major, err = handler.MajorUsecase.GetMajor(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message": "success get major by id",
			"data":   major,
		})
	}
}

func (handler MajorHandler) CreateMajor() echo.HandlerFunc {
	return func(e echo.Context) error {
		var major entity.Major
		if err := e.Bind(&major); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message": err.Error(),
			})
		}

		// Validasi input menggunakan package validator
		validate := validator.New()
		if err := validate.Struct(major); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message": err.Error(),
			})
		}
		err := handler.MajorUsecase.CreateMajor(major)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message": "failed to created major",
			})
		}
		return e.JSON(
			http.StatusCreated, map[string]interface{}{
			"status code": http.StatusCreated,
			"message": "success create new major",
			"data":   major,
		})
	}
}
func (handler MajorHandler) UpdateMajor() echo.HandlerFunc {
	var major entity.Major

	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message": err.Error(),
			})
		}

		err = handler.MajorUsecase.FindMajor(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		if err := e.Bind(&major); err != nil {
			return e.JSON(http.StatusNotFound, map[string]interface{}{
				"status code": http.StatusNotFound,
				"message": err.Error(),
			})
		}

		err = handler.MajorUsecase.UpdateMajor(id, major)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
				"status code": http.StatusOK,
				"message": "success update major",
				"data":major,

			})
	}
}

func (handler MajorHandler) DeleteMajor() echo.HandlerFunc {
	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message": "input id is not number",
			})
		}

		err = handler.MajorUsecase.DeleteMajor(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message": "Success Delete Major`",
		})
	}
}
