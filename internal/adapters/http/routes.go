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
	// educatoin news
	education_newsRepo    repository.EducationNewsRepository
	education_newsHandler handler.EducationNewsHandler
	education_newsUsecase usecase.EducationNewsUseCase
)

func declare() {
	userRepo = repository.UserRepository{DB: db.DbMysql}
	userUsecase = usecase.UserUseCase{Repo: userRepo}
	userHandler = handler.UserHandler{UserUsecase: userUsecase}
	// auth
	AuthHandler = handler.AuthHandler{Usecase: userUsecase}
	// education news
	education_newsRepo = repository.EducationNewsRepository{DB: db.DbMysql}
	education_newsUsecase = usecase.EducationNewsUseCase{Repo: education_newsRepo}
	education_newsHandler = handler.EducationNewsHandler{EducationNewsUsecase: education_newsUsecase}

}

func InitRoutes() *echo.Echo {
	db.Init()
	declare()

	e := echo.New()
	e.POST("/login", AuthHandler.Login())
	e.POST("/registrasi", AuthHandler.Register())

	// montor group
	mentors := e.Group("/mentors")
	mentors.Use(middleware.Logger())
	mentors.Use(middlewares.AuthMiddleware())
	mentors.Use(middlewares.RequireRole("mentors"))

	mentors.GET("/users", userHandler.GetAllUsers())
	mentors.GET("/users/:id", userHandler.GetUser())
	mentors.POST("/users", userHandler.CreateUser())
	mentors.DELETE("/users/:id", userHandler.DeleteUser())

	e.GET("/education_newses", education_newsHandler.GetAllEducationNewses())
	e.GET("/education_newses/:id", education_newsHandler.GetEducationNews())
	e.POST("/education_newses", education_newsHandler.CreateEducationNews())
	e.PUT("/education_newses/:id", education_newsHandler.UpdateEducationNews())
	e.DELETE("/education_newses/:id", education_newsHandler.DeleteEducationNews())

	// montor group
	students := e.Group("/students")
	students.Use(middleware.Logger())
	students.Use(middlewares.AuthMiddleware())
	students.Use(middlewares.RequireRole("students"))

	return e
}
