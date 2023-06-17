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
	// Section
	sectionRepo    repository.SectionRepository
	sectionHandler handler.SectionHandler
	sectionUsecase usecase.SectionUseCase
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
	//Promo
	promoRepo    repository.PromoRepository
	promoHandler handler.PromoHandler
	promoUsecase usecase.PromoUseCase
	//Rate course
	rateCourseRepo    repository.RateCourseRepository
	rateCourseHandler handler.RateCourseHandler
	rateCourseUsecase usecase.RateCourseUseCase

	//Trsaction
	transactionDetailRepo    repository.TrasanctionDetailsRepository
	transactionDetailUsecase usecase.TrasanctionDetailsUseCase

	//Trsaction
	transactionRepo    repository.TransactionRepository
	transactionHandler handler.TransactionHandler
	transactionUsecase usecase.TransactionUsecase
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
	// Section
	sectionRepo = repository.SectionRepository{DB: db.DbMysql}
	sectionUsecase = usecase.SectionUseCase{Repo: sectionRepo}
	sectionHandler = handler.SectionHandler{SectionUsecase: sectionUsecase}
	// Course
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
	// attachment
	attachmentRepo = repository.AttachmentRepository{DB: db.DbMysql}
	attachmentUsecase = usecase.AttachmentUseCase{Repo: attachmentRepo}
	attachmentHandler = handler.AttachmentHandler{AttachmentUsecase: attachmentUsecase}
	// Promo
	promoRepo = repository.PromoRepository{DB: db.DbMysql}
	promoUsecase = usecase.PromoUseCase{Repo: promoRepo}
	promoHandler = handler.PromoHandler{PromoUsecase: promoUsecase}
	// Promo
	rateCourseRepo = repository.RateCourseRepository{DB: db.DbMysql}
	rateCourseUsecase = usecase.RateCourseUseCase{Repo: rateCourseRepo}
	rateCourseHandler = handler.RateCourseHandler{RateCourseUsecase: rateCourseUsecase}

	// Promo
	transactionDetailRepo = repository.TrasanctionDetailsRepository{DB: db.DbMysql}
	transactionDetailUsecase = usecase.TrasanctionDetailsUseCase{TransactionDetailRepo: transactionDetailRepo}

	// Transaction
	transactionRepo = repository.TransactionRepository{DB: db.DbMysql}
	transactionUsecase = usecase.TransactionUsecase{TransactionRepo: transactionRepo, UserRepo: userRepo}

	transactionHandler = handler.TransactionHandler{
		TransactionUsecase:        transactionUsecase,
		Usecase:                   userUsecase,
		TrasanctionDetailsUseCase: transactionDetailUsecase}

}

func InitRoutes() *echo.Echo {
	db.Init()
	declare()

	e := echo.New()

	e.POST("/login", AuthHandler.Login())
	e.POST("/registrasi", AuthHandler.Register())
	e.POST("/registrasi-mentor", AuthHandler.MentorRegister())
	e.POST("/verify-otp", AuthHandler.VerifyOTP())
	e.POST("/forgot-password", AuthHandler.ForgotPassword())

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

	// route tas
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

	mentors.GET("/classes", classHandler.GetAllClasses())
	mentors.GET("/classes/:id", classHandler.GetClass())
	mentors.PUT("/classes/:id", classHandler.UpdateClass())
	mentors.POST("/classes", classHandler.CreateClass())
	mentors.DELETE("/classes/:id", classHandler.DeleteClass())

	mentors.GET("/categories", categoryHandler.GetAllCategories())
	mentors.GET("/categories/:id", categoryHandler.GetCategory())
	mentors.PUT("/cateories/:id", categoryHandler.UpdateCategory())
	mentors.POST("/categories", categoryHandler.CreateCategory())
	mentors.DELETE("/categories/:id", categoryHandler.DeleteCategory())

	mentors.GET("/majors", majorHandler.GetAllMajors())
	mentors.GET("/majors/:id", majorHandler.CreateMajor())
	mentors.PUT("/majors/:id", majorHandler.UpdateMajor())
	mentors.POST("/majors", majorHandler.CreateMajor())
	mentors.DELETE("/majors/:id", majorHandler.DeleteMajor())
	// route courses
	mentors.GET("/courses", courseHandler.GetAllCourses())
	mentors.GET("/courses/:id", courseHandler.GetCourse())
	mentors.PUT("/courses/:id", courseHandler.UpdateCourse())
	mentors.POST("/courses", courseHandler.CreateCourse())
	mentors.DELETE("/courses/:id", courseHandler.DeleteCourse())

	// route section
	mentors.GET("/section", sectionHandler.GetAllSections())
	mentors.GET("/section/:id", sectionHandler.CreateSection())
	mentors.PUT("/section/:id", sectionHandler.UpdateSection())
	mentors.POST("/section", sectionHandler.CreateSection())
	mentors.DELETE("/section/:id", sectionHandler.DeleteSection())
	// mentor logout
	mentors.POST("/logout", AuthHandler.Logout())

	mentors.GET("/classes", classHandler.GetAllClasses())
	mentors.GET("/classes/:id", classHandler.GetClass())
	mentors.GET("/class/filter", classHandler.FilterClasses())
	mentors.PUT("/classes/:id", classHandler.UpdateClass())
	mentors.POST("/classes", classHandler.CreateClass())
	mentors.DELETE("/classes/:id", classHandler.DeleteClass())

	mentors.GET("/categories", categoryHandler.GetAllCategories())
	mentors.GET("/categories/:id", categoryHandler.GetCategory())
	mentors.PUT("/cateories/:id", categoryHandler.UpdateCategory())
	mentors.POST("/categories", categoryHandler.CreateCategory())
	mentors.DELETE("/categories/:id", categoryHandler.DeleteCategory())

	mentors.GET("/majors", majorHandler.GetAllMajors())
	mentors.GET("/majors/:id", majorHandler.CreateMajor())
	mentors.GET("/majors/:id", majorHandler.GetMajor())
	mentors.GET("/majors/filter", majorHandler.FilterMajors())
	mentors.PUT("/majors/:id", majorHandler.UpdateMajor())
	mentors.POST("/majors", majorHandler.CreateMajor())
	mentors.DELETE("/majors/:id", majorHandler.DeleteMajor())

	mentors.GET("/courses", courseHandler.GetAllCourses())
	mentors.GET("/courses/:id", courseHandler.GetCourse())
	mentors.PUT("/courses/:id", courseHandler.UpdateCourse())
	mentors.POST("/courses", courseHandler.CreateCourse())
	mentors.DELETE("/courses/:id", courseHandler.DeleteCourse())
	mentors.GET("/courses/sort", courseHandler.GetAllCoursesSortedByField())
	mentors.GET("/promos", promoHandler.GetAllPromo())
	mentors.GET("/promos/:id", promoHandler.GetPromo())
	mentors.PUT("/promos/:id", promoHandler.UpdatePromo())
	mentors.POST("/promos", promoHandler.CreatePromo())
	mentors.DELETE("/promos/:id", promoHandler.DeletePromo())



	// students group
	students := e.Group("/students")
	students.Use(middleware.Logger())
	students.Use(middlewares.AuthMiddleware())
	students.Use(middlewares.RequireRole("students"))

	//route course student
	students.GET("/courses/:userID", courseHandler.GetCoursesByUserID)
	students.GET("/courses/status", courseHandler.GetCoursesStatus)
	students.GET("/courses/module", courseHandler.GetAllModules())
	students.GET("/courses/module/:id", courseHandler.GetModule())

	students.POST("/new-password", AuthHandler.NewPassword())
	students.GET("/courses/:userID", courseHandler.GetCoursesByUserID)
	students.GET("/courses/status", courseHandler.GetCoursesStatus)

	// route attachment
	students.GET("/attachment/:id", attachmentHandler.GetAllAttachments())
	students.GET("/attachment/find/:id", attachmentHandler.GetAttachment())
	students.POST("/attachment", attachmentHandler.CreateAttachment())
	students.DELETE("/attachment/:id", attachmentHandler.DeleteAttachment())
	students.GET("/courses/quiz-attachments", attachmentHandler.GetQuizAttachments)
	students.GET("/courses/quiz-attachments/:id", attachmentHandler.GetQuizAttachmentByID)
	students.GET("/courses/materi-attachments", attachmentHandler.GetMateriAttachments)
	students.GET("/courses/materi-attachments/:id", attachmentHandler.GetMateriAttachmentByID)

	students.GET("/classes", classHandler.GetAllClasses())
	students.GET("/classes/:id", classHandler.GetClass())
	students.GET("/categories", categoryHandler.GetAllCategories())
	students.GET("/categories/:id", categoryHandler.GetCategory())
	students.GET("/majors", majorHandler.GetAllMajors())
	students.GET("/majors/:id", majorHandler.GetMajor())
	students.GET("/courses", courseHandler.GetAllCourseStudents())
	students.GET("/courses/:id", courseHandler.GetCourse())
	students.GET("/promos", promoHandler.GetAllPromo())
	students.GET("/promos/:id", promoHandler.GetPromo())
	students.GET("/class/filter", classHandler.FilterClasses())
	students.GET("/majors/filter", majorHandler.FilterMajors())
	// chat mentor
	students.GET("/mentors", userHandler.GetUserByRole())
	students.GET("/mentors/:id", userHandler.GetUser())
	// rate course
	students.POST("/rate-course", rateCourseHandler.CreateRateCourse())
	students.GET("/courses/sort", courseHandler.GetAllCoursesSortedByField())
	students.PUT("/user/profile", userHandler.UpdateUser())
	// transaction
	students.POST("/transaction", transactionHandler.CheckoutTransaction())

	return e
}
