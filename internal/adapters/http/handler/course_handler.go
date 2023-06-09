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

type CourseHandler struct {
	CourseUsecase usecase.CourseUseCase
}

func (handler CourseHandler) GetAllCourses() echo.HandlerFunc {
	return func(e echo.Context) error {
		var courses []entity.Course
		MentorId, err := service.GetUserIDFromToken(e)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		courses, err = handler.CourseUsecase.GetAllCourses(MentorId)
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

func (handler CourseHandler) GetAllCourseStudents() echo.HandlerFunc {
	return func(e echo.Context) error {
		var courses []entity.Course

		courses, err := handler.CourseUsecase.GetAllCourseStudents()
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
		var countSection []entity.CourseWithSectionCount

		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     err.Error(),
			})
		}
		countSection, err = handler.CourseUsecase.GetAllCoursesWithSectionCount(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		data := make(map[string]interface{})
		data["coursesCount"] = countSection

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "success get course by id",
			"data":        data,
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
				"message":     err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "success get course by mentor id",
			"data":        course,
		})
	}
}
func (handler CourseHandler) CreateCourse() echo.HandlerFunc {
	return func(e echo.Context) error {
		var course entity.Course
		MentorId, err := service.GetUserIDFromToken(e)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		course.MentorId = MentorId
		course.Status = "draft"

		if err := e.Bind(&course); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     err.Error(),
			})
		}

		validate := validator.New()
		if err := validate.Struct(course); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     err.Error(),
			})
		}
		if err := handler.CourseUsecase.CreateCourse(&course); err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     "failed to created course",
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

func (handler CourseHandler) UpdateCourse() echo.HandlerFunc {
	var course entity.Course

	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     err.Error(),
			})
		}

		err = handler.CourseUsecase.FindCourse(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		if err := e.Bind(&course); err != nil {
			return e.JSON(http.StatusNotFound, map[string]interface{}{
				"status code": http.StatusNotFound,
				"message":     err.Error(),
			})
		}
		MentorId, err := service.GetUserIDFromToken(e)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		course.MentorId = MentorId

		err = handler.CourseUsecase.UpdateCourse(id, course)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "success update course",
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

func (handler CourseHandler) GetCoursesStatus(c echo.Context) error {
	userID, err := service.GetUserIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status code": http.StatusInternalServerError,
			"message":     err.Error(),
		})
	}
	coursesStatus, err := handler.CourseUsecase.GetCoursesStatus(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status code": http.StatusInternalServerError,
			"message":     err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status code": http.StatusOK,
		"data":        coursesStatus,
	})
}

func (handler CourseHandler) GetAllModules() echo.HandlerFunc {
	return func(e echo.Context) error {
		var modules []entity.Module

		modules, err := handler.CourseUsecase.GetAllModules()
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "success get all modules",
			"data":        modules,
		})
	}
}

func (handler CourseHandler) GetModule() echo.HandlerFunc {
	return func(e echo.Context) error {
		var module entity.Module
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     err.Error(),
			})
		}

		module, err = handler.CourseUsecase.GetModule(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "success get module by id",
			"data":        module,
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

func (handler CourseHandler) GetCourseSection() echo.HandlerFunc {
	return func(e echo.Context) error {
		var course entity.Course

		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     err.Error(),
			})
		}
		course, err = handler.CourseUsecase.GetCourseSection(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		data := make(map[string]interface{})
		data["courses"] = course

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "success get course by id",
			"data":        data,
		})
	}
}

func (handler CourseHandler) GetStudentsByCourseID(c echo.Context) error {
	courseID, err := strconv.Atoi(c.Param("course_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status code": http.StatusBadRequest,
			"message":     "Invalid course ID",
		})
	}

	users, err := handler.CourseUsecase.GetStudentsByCourseID(courseID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status code": http.StatusInternalServerError,
			"message":     err.Error(),
		})
	}

	response := map[string]interface{}{
		"status code": http.StatusOK,
		"message":     "Success get users by course ID",
		"data":        users,
	}

	return c.JSON(http.StatusOK, response)
}

func (handler CourseHandler) GetAllCoursesWithSectionAndStudentCount() echo.HandlerFunc {
	return func(e echo.Context) error {
		mentorId, err := service.GetUserIDFromToken(e)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}
		courses, err := handler.CourseUsecase.GetAllCoursesWithSectionAndStudentCount(mentorId)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "success get course by mentor id",
			"data":        courses,
		})
	}
}
