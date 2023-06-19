package handler

import (
	"net/http"
	"strconv"

	"capston-lms/internal/application/usecase"
	"capston-lms/internal/entity"

	// "github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type EducationNewsHandler struct {
	EducationNewsUsecase usecase.EducationNewsUseCase
}

func (handler EducationNewsHandler) GetAllEducationNewses() echo.HandlerFunc {
	return func(e echo.Context) error {
		var education_newses []entity.EducationNews

		education_newses, err := handler.EducationNewsUsecase.GetAllEducationNewses()
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "success get all news",
			"data":        education_newses,
		})
	}
}

func (handler EducationNewsHandler) GetEducationNews() echo.HandlerFunc {
	return func(e echo.Context) error {
		var education_news entity.EducationNews
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     err.Error(),
			})
		}

		education_news, err = handler.EducationNewsUsecase.GetEducationNews(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "success get news by id",
			"data":        education_news,
		})
	}
}

func (handler EducationNewsHandler) CreateEducationNews() echo.HandlerFunc {
	return func(e echo.Context) error {
		var education_news entity.EducationNews
		if err := e.Bind(&education_news); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     err.Error(),
			})
		}

		// validate := validator.New()
		// if err := validate.Struct(education_news); err != nil {
		// 	return e.JSON(http.StatusBadRequest, map[string]interface{}{
		// 		"status code": http.StatusBadRequest,
		// 		"message":     err.Error(),
		// 	})
		// }

		err := handler.EducationNewsUsecase.CreateEducationNews(education_news)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     "failed to created user",
			})
		}

		return e.JSON(
			http.StatusCreated, map[string]interface{}{
				"status code": http.StatusCreated,
				"message":     "success create new news",
				"data":        education_news,
			})
	}
}

func (handler EducationNewsHandler) UpdateEducationNews() echo.HandlerFunc {
	var education_news entity.EducationNews

	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "input id is not a number",
			})
		}

		err = handler.EducationNewsUsecase.FindEducationNews(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"message": "Record Not Found",
			})
		}

		if err := e.Bind(&education_news); err != nil {
			return e.JSON(400, echo.Map{
				"error": err.Error(),
			})
		}

		err = handler.EducationNewsUsecase.UpdateEducationNews(id, education_news)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "success get news by id",
			"data":        education_news,
		})
	}
}

func (handler EducationNewsHandler) DeleteEducationNews() echo.HandlerFunc {
	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     "input id is not number",
			})
		}

		err = handler.EducationNewsUsecase.DeleteEducationNews(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "Success Delete News`",
		})
	}
}
