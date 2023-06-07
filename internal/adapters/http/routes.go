package http

import (
	db "capston-lms/internal/adapters/db/mysql"
	handler "capston-lms/internal/adapters/http/handler"
	middlewares "capston-lms/internal/adapters/http/middleware"
	repository "capston-lms/internal/adapters/repository"
	usecase "capston-lms/internal/application/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	// user management
	userRepo    repository.UserRepository
	userHandler handler.UserHandler
	userUsecase usecase.UserUseCase
	// auth
	AuthHandler handler.AuthHandler

	// class
	classRepo    repository.ClassRepository
	classHandler handler.ClassHandler
	classUsecase usecase.ClassUseCase
	// Category
	categoryRepo    repository.CategoryRepository
	categoryHandler handler.CategoryHandler
	categoryUsecase usecase.CategoryUseCase
	// Major
	majorRepo    repository.MajorRepository
	majorHandler handler.MajorHandler
	majorUsecase usecase.MajorUseCase
	// Course
	courseEnrollmentRepo    repository.CourseEnrollmentRepository
	courseEnrollmentHandler handler.CourseEnrollmentHandler
	courseEnrollmentUseCase usecase.CourseEnrollmentUseCase

	// course enrollment
	courseRepo    repository.CourseRepository
	courseHandler handler.CourseHandler
	courseUsecase usecase.CourseUseCase

	// attachment
	attachmentRepo    repository.AttachmentRepository
	attachmentHandler handler.AttachmentHandler
	attachmentUsecase usecase.AttachmentUseCase
)

func declare() {
	// user
	userRepo = repository.UserRepository{DB: db.DbMysql}
	userUsecase = usecase.UserUseCase{Repo: userRepo}
	userHandler = handler.UserHandler{UserUsecase: userUsecase}
	// auth
	AuthHandler = handler.AuthHandler{Usecase: userUsecase}

	// class
	classRepo = repository.ClassRepository{DB: db.DbMysql}
	classUsecase = usecase.ClassUseCase{Repo: classRepo}
	classHandler = handler.ClassHandler{ClassUsecase: classUsecase}
	// category
	categoryRepo = repository.CategoryRepository{DB: db.DbMysql}
	categoryUsecase = usecase.CategoryUseCase{Repo: categoryRepo}
	categoryHandler = handler.CategoryHandler{CategoryUsecase: categoryUsecase}
	// Major
	majorRepo = repository.MajorRepository{DB: db.DbMysql}
	majorUsecase = usecase.MajorUseCase{Repo: majorRepo}
	majorHandler = handler.MajorHandler{MajorUsecase: majorUsecase}
	// Major
	courseRepo = repository.CourseRepository{DB: db.DbMysql}
	courseUsecase = usecase.CourseUseCase{Repo: courseRepo}
	courseHandler = handler.CourseHandler{CourseUsecase: courseUsecase}

	// course enrrolment
	courseEnrollmentRepo = repository.CourseEnrollmentRepository{DB: db.DbMysql}
	courseEnrollmentUseCase = usecase.CourseEnrollmentUseCase{CourseEnrollmentRepo: courseEnrollmentRepo}
	courseEnrollmentHandler = handler.CourseEnrollmentHandler{CourseEnrollmentUseCase: courseEnrollmentUseCase}

	// attachment
	attachmentRepo = repository.AttachmentRepository{DB: db.DbMysql}
	attachmentUsecase = usecase.AttachmentUseCase{Repo: attachmentRepo}
	attachmentHandler = handler.AttachmentHandler{AttachmentUsecase: attachmentUsecase}
}

func InitRoutes() *echo.Echo {
	db.Init()
	declare()

	// Middleware untuk mengizinkan header "Content-Type: application/json"
	jsonMiddleware := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			return next(c)
		}
	}

	e := echo.New()
	e.POST("/login", AuthHandler.Login())
	e.POST("/registrasi", AuthHandler.Register())
	e.POST("/verify-otp", AuthHandler.VerifyOTP(), jsonMiddleware)

	// montor group
	mentors := e.Group("/mentors")
	mentors.Use(middleware.Logger())
	mentors.Use(middlewares.AuthMiddleware())
	mentors.Use(middlewares.RequireRole("mentors"))

	mentors.GET("/users", userHandler.GetAllUsers())
	mentors.GET("/users/:id", userHandler.GetUser())
	mentors.POST("/users", userHandler.CreateUser())
	mentors.DELETE("/users/:id", userHandler.DeleteUser())

	mentors.GET("/chat/students/:id", courseEnrollmentHandler.GetAllStudents())
	mentors.GET("/chat/courses", courseEnrollmentHandler.GetAllCourse())

	e.GET("/classes", classHandler.GetAllClasses())
	e.GET("/classes/:id", classHandler.GetClass())
	e.POST("/classes", classHandler.CreateClass())
	e.DELETE("/classes/:id", classHandler.DeleteClass())

	e.GET("/categories", categoryHandler.GetAllCategories())
	e.GET("/categories/:id", categoryHandler.GetCategory())
	e.POST("/categories", categoryHandler.CreateCategory())
	e.DELETE("/categories/:id", categoryHandler.DeleteCategory())

	e.GET("/majors", majorHandler.GetAllMajors())
	e.GET("/majors/:id", majorHandler.CreateMajor())
	e.POST("/majors", majorHandler.CreateMajor())
	e.DELETE("/majors/:id", majorHandler.DeleteMajor())

	e.GET("/courses", courseHandler.GetAllCourses())
	e.GET("/courses/:id", courseHandler.CreateCourse())
	e.POST("/courses", courseHandler.CreateCourse())
	e.DELETE("/courses/:id", courseHandler.DeleteCourse())

	// students group
	students := e.Group("/students")
	students.Use(middleware.Logger())
	students.Use(middlewares.AuthMiddleware())
	students.Use(middlewares.RequireRole("students"))
	students.GET("/courses/:userID", courseHandler.GetCoursesByUserID)

	// route attachment
	students.GET("/attachment/:id", attachmentHandler.GetAllAttachments())
	students.GET("/attachment/find/:id", attachmentHandler.GetAttachment())
	students.POST("/attachment", attachmentHandler.CreateAttachment())
	students.DELETE("/attachment/:id", attachmentHandler.DeleteAttachment())

	return e
}

// Middleware untuk mengizinkan header Content-Type: application/json
func allowJSONContentType(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")

		contentType := c.Request().Header.Get("Content-Type")
		if contentType != "application/json" {
			return echo.NewHTTPError(http.StatusUnsupportedMediaType, "Only application/json content type is allowed")
		}

		return next(c)
	}
}
