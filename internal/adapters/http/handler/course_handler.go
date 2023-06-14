package handler

import (
	"net/http"
	"strconv"

	"capston-lms/internal/application/service"
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

func (handler CourseHandler) GetCourseByMentorId() echo.HandlerFunc {
	return func(e echo.Context) error {
		var course entity.Course
		mentorId, err := service.GetUserIDFromToken(e)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}
		course, err = handler.CourseUsecase.GetCourseByMentorId(mentorId)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message": "success get course by mentor id",
			"data":   course,
		})
	}
}
func (handler CourseHandler) CreateCourse() echo.HandlerFunc {
	return func(e echo.Context) error {
		var course entity.Course
		mentorId, err := service.GetUserIDFromToken(e)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}
		course.MentorId=mentorId
		if err := e.Bind(&course); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message": err.Error(),
			})
		}
		// Validasi input menggunakan package validator
		validate := validator.New()
		if err := validate.Struct(course); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message": err.Error(),
			})
		}
		if err := handler.CourseUsecase.CreateCourse(course); err != nil {
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

func (handler CourseHandler) GetAllCoursesSortedByField() echo.HandlerFunc {
	return func(e echo.Context) error {
		sortBy := e.QueryParam("sort_by")
		ascending, _ := strconv.ParseBool(e.QueryParam("ascending"))

		var courses []entity.Course
		var err error

		switch sortBy {
		case "Segera Selesai":
			courses, err = handler.CourseUsecase.GetAllCoursesSortedByCompletion(ascending)
		case "Kursus Baru":
			courses, err = handler.CourseUsecase.GetAllCoursesSortedByNewness(ascending)
		case "Kelas Atas":
			courses, err = handler.CourseUsecase.GetAllCoursesSortedByHighLevel(ascending)
		case "Kelas Bawah":
			courses, err = handler.CourseUsecase.GetAllCoursesSortedByLowLevel(ascending)
		default:
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     "Invalid sort_by parameter",
			})
		}

		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "Success get all courses sorted by field",
			"data":        courses,
		})
	}
}
