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
	// folder
	folderRepo    repository.FolderRepository
	folderHandler handler.FolderHandler
	folderUsecase usecase.FolderUseCase
	// attachment
	attachmentRepo    repository.AttachmentRepository
	attachmentHandler handler.AttachmentHandler
	attachmentUsecase usecase.AttachmentUseCase
	// module
	moduleRepo    repository.ModuleRepository
	moduleHandler handler.ModuleHandler
	moduleUsecase usecase.ModuleUseCase
	// task
	taskRepo    repository.TaskRepository
	taskHandler handler.TaskHandler
	taskUsecase usecase.TaskUseCase
	// submission
	submissionRepo    repository.SubmissionRepository
	submissionHandler handler.SubmissionHandler
	submissionUsecase usecase.SubmissionUseCase
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
	// folder
	folderRepo = repository.FolderRepository{DB: db.DbMysql}
	folderUsecase = usecase.FolderUseCase{Repo: folderRepo}
	folderHandler = handler.FolderHandler{FolderUsecase: folderUsecase}
	// attachments
	attachmentRepo = repository.AttachmentRepository{DB: db.DbMysql}
	attachmentUsecase = usecase.AttachmentUseCase{Repo: attachmentRepo}
	attachmentHandler = handler.AttachmentHandler{AttachmentUsecase: attachmentUsecase}
	// attachments
	moduleRepo = repository.ModuleRepository{DB: db.DbMysql}
	moduleUsecase = usecase.ModuleUseCase{Repo: moduleRepo}
	moduleHandler = handler.ModuleHandler{ModuleUseCase: moduleUsecase}
	// task
	taskRepo = repository.TaskRepository{DB: db.DbMysql}
	taskUsecase = usecase.TaskUseCase{Repo: taskRepo}
	taskHandler = handler.TaskHandler{TaskUseCase: taskUsecase}
	// task
	submissionRepo = repository.SubmissionRepository{DB: db.DbMysql}
	submissionUsecase = usecase.SubmissionUseCase{Repo: submissionRepo}
	submissionHandler = handler.SubmissionHandler{SubmissionUseCase: submissionUsecase}
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

	mentors.GET("/chat/students/:id", courseEnrollmentHandler.GetAllStudents())
	mentors.GET("/chat/courses", courseEnrollmentHandler.GetAllCourse())
	// route folders
	mentors.GET("/folders", folderHandler.GetAllFolders())
	mentors.GET("/folders/:id", folderHandler.GetFolder())
	mentors.POST("/folders", folderHandler.CreateFolder())
	mentors.DELETE("/folders/:id", folderHandler.DeleteFolder())

	// route attachment
	mentors.GET("/attachment/:id", attachmentHandler.GetAllAttachments())
	mentors.GET("/attachment/find/:id", attachmentHandler.GetAttachment())
	mentors.POST("/attachment", attachmentHandler.CreateAttachment())
	mentors.DELETE("/attachment/:id", attachmentHandler.DeleteAttachment())

	// route modules
	mentors.GET("/module", moduleHandler.GetAllModules())
	mentors.GET("/module/:id", moduleHandler.GetModule())
	mentors.PUT("/module/:id", moduleHandler.UpdateModule())
	mentors.POST("/module", moduleHandler.CreateModule())
	mentors.DELETE("/module/:id", moduleHandler.DeleteModule())

	// route modules
	mentors.GET("/task", taskHandler.GetAllTasks())
	mentors.GET("/task/:id", taskHandler.GetTask())
	mentors.PUT("/task/:id", taskHandler.UpdateTask())
	mentors.POST("/task", taskHandler.CreateTask())
	mentors.DELETE("/task/:id", taskHandler.DeleteTask())
	// route submissions
	mentors.GET("/submission", submissionHandler.GetAllSubmissions())
	mentors.GET("/submission/:id", submissionHandler.GetSubmission())
	mentors.PUT("/submission/:id", submissionHandler.UpdateSubmission())
	mentors.POST("/submission", submissionHandler.CreateSubmission())
	mentors.DELETE("/submission/:id", submissionHandler.DeleteSubmission())

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

	return e
}
