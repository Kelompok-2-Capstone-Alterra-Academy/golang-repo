package handler

import (
	"net/http"
	"strconv"

	"capston-lms/internal/application/usecase"
	"capston-lms/internal/entity"

	// "github.com/golang-jwt/jwt"

	"github.com/go-playground/validator"
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
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message": "success get all course",
			"data":   courses,
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
				"message": err.Error(),
			})
		}

		course, err = handler.CourseUsecase.GetCourse(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message": "success get course by id",
			"data":   course,
		})
	}
}

func (handler CourseHandler) CreateCourse() echo.HandlerFunc {
	return func(e echo.Context) error {
		var course entity.Course
		
		if err := e.Bind(&course); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message": err.Error(),
			})
		}
		// user := e.Get("user").(*jwt.Token)
		// claims := user.Claims.(*jwt.MapClaims)
		// mentorID := int((*claims)["id"].(int))
		// course.MentorId = int(mentorID)
		// Validasi input menggunakan package validator
		validate := validator.New()
		if err := validate.Struct(course); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message": err.Error(),
			})
		}
		err := handler.CourseUsecase.CreateCourse(course)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message": "failed to created course",
			})
		}
		return e.JSON(
			http.StatusCreated, map[string]interface{}{
			"status code": http.StatusCreated,
			"message": "success create new course",
			"data":   course,
		})
	}
}

func (handler CourseHandler) UpdateCourse() echo.HandlerFunc {
	var course entity.Course

	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message": err.Error(),
			})
		}

		err = handler.CourseUsecase.FindCourse(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		if err := e.Bind(&course); err != nil {
			return e.JSON(http.StatusNotFound, map[string]interface{}{
				"status code": http.StatusNotFound,
				"message": err.Error(),
			})
		}

		err = handler.CourseUsecase.UpdateCourse(id, course)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
				"status code": http.StatusOK,
				"message": "success update course",
				"data":course,

			})
	}
}
func (handler CourseHandler) DeleteCourse() echo.HandlerFunc {
	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message": "input id is not number",
			})
		}

		err = handler.CourseUsecase.DeleteCourse(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message": "Success Delete Course`",
		})
	}
}
