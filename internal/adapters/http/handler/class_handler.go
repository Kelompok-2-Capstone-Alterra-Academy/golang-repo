package handler

import (
	"net/http"
	"strconv"

	"capston-lms/internal/application/usecase"
	"capston-lms/internal/entity"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type ClassHandler struct {
	ClassUsecase usecase.ClassUseCase
}

func (handler ClassHandler) GetAllClasses() echo.HandlerFunc {
	return func(e echo.Context) error {
		var classes []entity.Class

		classes, err := handler.ClassUsecase.GetAllClasses()
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message": "success get all class",
			"data":   classes,
		})
	}
}

func (handler ClassHandler) GetClass() echo.HandlerFunc {
	return func(e echo.Context) error {
		var class entity.Class
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message": err.Error(),
			})
		}

		class, err = handler.ClassUsecase.GetClass(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message": "success get class by id",
			"data":   class,
		})
	}
}

func (handler ClassHandler) FilterClasses() echo.HandlerFunc {
	return func(c echo.Context) error {
		var classes []entity.Class
		selectedClass := c.QueryParams()["class"]
		filteredClass := make([]entity.Class, 0)

		for _, class := range classes {
			for _, selectedClass := range selectedClass {
				if class.ClassName == selectedClass {
					filteredClass = append(filteredClass, class)
				}
			}
		}

		return c.JSON(http.StatusOK, filteredClass)
	}
}

func (handler ClassHandler) CreateClass() echo.HandlerFunc {
	return func(e echo.Context) error {
		var class entity.Class
		if err := e.Bind(&class); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message": err.Error(),
			})
		}

		// Validasi input menggunakan package validator
		validate := validator.New()
		if err := validate.Struct(class); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message": err.Error(),
			})
		}
		err := handler.ClassUsecase.CreateClass(class)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message": "failed to created class",
			})
		}
		return e.JSON(
			http.StatusCreated, map[string]interface{}{
			"status code": http.StatusCreated,
			"message": "success create new class",
			"data":   class,
		})
	}
}

func (handler ClassHandler) UpdateClass() echo.HandlerFunc {
	var class entity.Class

	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message": err.Error(),
			})
		}

		err = handler.ClassUsecase.FindClass(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		if err := e.Bind(&class); err != nil {
			return e.JSON(http.StatusNotFound, map[string]interface{}{
				"status code": http.StatusNotFound,
				"message": err.Error(),
			})
		}

		err = handler.ClassUsecase.UpdateClass(id, class)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
				"status code": http.StatusOK,
				"message": "success update class",
				"data":class,

			})
	}
}
func (handler ClassHandler) DeleteClass() echo.HandlerFunc {
	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message": "input id is not number",
			})
		}

		err = handler.ClassUsecase.DeleteClass(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message": "Success Delete Class`",
		})
	}
}
