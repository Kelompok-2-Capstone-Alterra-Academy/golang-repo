package handler

import (
	"net/http"
	"strconv"

	"capston-lms/internal/application/service"
	"capston-lms/internal/application/usecase"
	"capston-lms/internal/entity"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type RateCourseHandler struct {
	RateCourseUsecase usecase.RateCourseUseCase
}

func (handler RateCourseHandler) CreateRateCourse() echo.HandlerFunc {
	return func(e echo.Context) error {
		var rateCourse entity.RateCourse
		userID, err := service.GetUserIDFromToken(e)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     "failed to get user ID",
			})
		}
		rateCourse.UserId = userID

		courseID := e.Param("course_id")
		courseIDInt, err := strconv.Atoi(courseID)
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     "invalid course ID",
			})
		}
		rateCourse.CourseId = courseIDInt

		if err := e.Bind(&rateCourse); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     err.Error(),
			})
		}
		// Validasi input menggunakan package validator
		validate := validator.New()
		if err := validate.Struct(rateCourse); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     err.Error(),
			})
		}

		err = handler.RateCourseUsecase.CreateRateCourse(rateCourse)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     "failed to create rateCourse",
			})
		}

		return e.JSON(http.StatusCreated, map[string]interface{}{
			"status code": http.StatusCreated,
			"message":     "success create new rateCourse",
			"data":        rateCourse,
		})
	}
}
