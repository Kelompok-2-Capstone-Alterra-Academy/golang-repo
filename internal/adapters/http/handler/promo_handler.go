package handler

import (
	"net/http"
	"strconv"

	"capston-lms/internal/application/usecase"
	"capston-lms/internal/entity"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type PromoHandler struct {
	PromoUsecase usecase.PromoUseCase
}

func (handler PromoHandler) GetAllPromo() echo.HandlerFunc {
	return func(e echo.Context) error {
		var promos []entity.Promo

		promos, err := handler.PromoUsecase.GetAllPromo()
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message": "success get all promo",
			"data":   promos,
		})
	}
}

func (handler PromoHandler) GetPromo() echo.HandlerFunc {
	return func(e echo.Context) error {
		var promo entity.Promo
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message": err.Error(),
			})
		}

		promo, err = handler.PromoUsecase.GetPromo(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message": "success get promo by id",
			"data":   promo,
		})
	}
}

func (handler PromoHandler) CreatePromo() echo.HandlerFunc {
	return func(e echo.Context) error {
		var promo entity.Promo
		if err := e.Bind(&promo); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message": err.Error(),
			})
		}

		// Validasi input menggunakan package validator
		validate := validator.New()
		if err := validate.Struct(promo); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message": err.Error(),
			})
		}
		err := handler.PromoUsecase.CreatePromo(promo)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message": "failed to create promo",
			})
		}
		return e.JSON(
			http.StatusCreated, map[string]interface{}{
			"status code": http.StatusCreated,
			"message": "success create new promo",
			"data":   promo,
		})
	}
}

func (handler PromoHandler) UpdatePromo() echo.HandlerFunc {
	var promo entity.Promo

	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message": err.Error(),
			})
		}

		err = handler.PromoUsecase.FindPromo(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		if err := e.Bind(&promo); err != nil {
			return e.JSON(http.StatusNotFound, map[string]interface{}{
				"status code": http.StatusNotFound,
				"message": err.Error(),
			})
		}

		err = handler.PromoUsecase.UpdatePromo(id, promo)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
				"status code": http.StatusOK,
				"message": "success update promo",
				"data":promo,

			})
	}
}

func (handler PromoHandler) DeletePromo() echo.HandlerFunc {
	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message": "input id is not number",
			})
		}

		err = handler.PromoUsecase.DeletePromo(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message": "Success Delete Promo`",
		})
	}
}
