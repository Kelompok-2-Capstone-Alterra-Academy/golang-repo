package http

import (
	db "capston-lms/internal/adapters/db/mysql"
	handler "capston-lms/internal/adapters/http/handler"
	middlewares "capston-lms/internal/adapters/http/middleware"
	repository "capston-lms/internal/adapters/repository"
	usecase "capston-lms/internal/application/usecase"

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
)

func declare() {
	userRepo = repository.UserRepository{DB: db.DbMysql}
	userUsecase = usecase.UserUseCase{Repo: userRepo}
	userHandler = handler.UserHandler{UserUsecase: userUsecase}
	// auth
	AuthHandler = handler.AuthHandler{Usecase: userUsecase}

}

func InitRoutes() *echo.Echo {
	db.Init()
	declare()

	e := echo.New()
	e.POST("/login", AuthHandler.Login())
	e.POST("/registrasi", AuthHandler.Register())
	e.POST("/verify-otp", AuthHandler.VerifyOTP())

	// montor group
	mentors := e.Group("/mentors")
	mentors.Use(middleware.Logger())
	mentors.Use(middlewares.AuthMiddleware())
	mentors.Use(middlewares.RequireRole("mentors"))

	mentors.GET("/users", userHandler.GetAllUsers())
	mentors.GET("/users/:id", userHandler.GetUser())
	mentors.POST("/users", userHandler.CreateUser())
	mentors.DELETE("/users/:id", userHandler.DeleteUser())

	// montor group
	students := e.Group("/students")
	students.Use(middleware.Logger())
	students.Use(middlewares.AuthMiddleware())
	students.Use(middlewares.RequireRole("students"))

	return e
}
