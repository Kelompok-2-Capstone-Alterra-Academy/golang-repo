package handler

import (
	"net/http"
	"strconv"

	"capston-lms/internal/application/service"
	"capston-lms/internal/application/usecase"
	"capston-lms/internal/entity"

	"github.com/labstack/echo/v4"
)

type CourseHandler struct {
	CourseUsecase usecase.CourseUseCase
}

func (handler CourseHandler) GetAllCourses() echo.HandlerFunc {
	return func(e echo.Context) error {
		var courses []entity.Course

		courses, err := handler.CourseUsecase.GetAllCourses()
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "success get all course",
			"data":        courses,
		})
	}
}

func (handler CourseHandler) GetCourse() echo.HandlerFunc {
	return func(e echo.Context) error {
		var course entity.Course
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     err.Error(),
			})
		}

		course, err = handler.CourseUsecase.GetCourse(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "success get course by id",
			"data":        course,
		})
	}
}

func (handler CourseHandler) CreateCourse() echo.HandlerFunc {
	return func(e echo.Context) error {
		var course entity.Course
		if err := e.Bind(&course); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     err.Error(),
			})
		}

		// // Validasi input menggunakan package validator
		// validate := validator.New()
		// if err := validate.Struct(course); err != nil {
		// 	return e.JSON(http.StatusBadRequest, map[string]interface{}{
		// 		"status code": http.StatusBadRequest,
		// 		"message": err.Error(),
		// 	})
		// }

		err := handler.CourseUsecase.CreateCourse(course)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     "failed to created user",
			})
		}

		return e.JSON(
			http.StatusCreated, map[string]interface{}{
				"status code": http.StatusCreated,
				"message":     "success create new course",
				"data":        course,
			})
	}
}

func (handler CourseHandler) DeleteCourse() echo.HandlerFunc {
	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     "input id is not number",
			})
		}

		err = handler.CourseUsecase.DeleteCourse(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "Success Delete Course`",
		})
	}
}

func (handler CourseHandler) GetCoursesByUserID(c echo.Context) error {
	userID, err := service.GetUserIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status code": http.StatusInternalServerError,
			"message":     err.Error(),
		})
	}

	courses, err := handler.CourseUsecase.GetCoursesByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status code": http.StatusInternalServerError,
			"message":     err.Error(),
		})
	}

	response := map[string]interface{}{
		"status code": http.StatusOK,
		"message":     "Success get course by user ID and course ID",
		"data":        courses,
	}

	return c.JSON(http.StatusOK, response)
}
