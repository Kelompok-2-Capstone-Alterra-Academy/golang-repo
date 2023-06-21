package usecase

import (
	"capston-lms/internal/adapters/repository"
	"capston-lms/internal/entity"
)

type CourseUseCase struct {
	Repo                 repository.CourseRepository
	CourseEnrollmentRepo repository.CourseEnrollmentRepository
}

func (usecase CourseUseCase) GetAllCourses(mentorId int) ([]entity.Course, error) {
	courses, err := usecase.Repo.GetAllCourses(mentorId)
	return courses, err
}

func (usecase CourseUseCase) GetAllCourseStudents() ([]entity.Course, error) {
	courses, err := usecase.Repo.GetAllCourseStudents()
	return courses, err
}

func (usecase CourseUseCase) GetCourse(id int) (entity.Course, error) {
	course, err := usecase.Repo.GetCourse(id)
	return course, err
}

func (usecase CourseUseCase) GetCourseSection(id int) (entity.Course, error) {
	course, err := usecase.Repo.GetCourseSection(id)
	return course, err
}
func (usecase CourseUseCase) GetCourseByMentorId(id int) (entity.Course, error) {
	course, err := usecase.Repo.GetCourseByMentorId(id)
	return course, err
}

func (usecase CourseUseCase) CreateCourse(course *entity.Course) error {
	err := usecase.Repo.CreateCourse(course)
	return err
}

func (usecase CourseUseCase) UpdateCourse(id int, course entity.Course) error {
	err := usecase.Repo.UpdateCourse(id, course)
	return err
}

func (usecase CourseUseCase) DeleteCourse(id int) error {
	err := usecase.Repo.DeleteCourse(id)
	return err
}
func (usecase CourseUseCase) FindCourse(id int) error {
	err := usecase.Repo.FindCourse(id)
	return err
}

func (usecase CourseUseCase) GetAllCoursesSortedByCompletion(ascending bool) ([]entity.Course, error) {
	courses, err := usecase.Repo.GetAllCoursesSortedByCompletion(ascending)
	return courses, err
}

func (usecase CourseUseCase) GetAllCoursesSortedByNewness(ascending bool) ([]entity.Course, error) {
	courses, err := usecase.Repo.GetAllCoursesSortedByNewness(ascending)
	return courses, err
}

func (usecase CourseUseCase) GetAllCoursesSortedByHighLevel(ascending bool) ([]entity.Course, error) {
	courses, err := usecase.Repo.GetAllCoursesSortedByHighLevel(ascending)
	return courses, err
}

func (usecase CourseUseCase) GetAllCoursesSortedByLowLevel(ascending bool) ([]entity.Course, error) {
	courses, err := usecase.Repo.GetAllCoursesSortedByLowLevel(ascending)
	return courses, err
}

func (usecase CourseUseCase) GetCoursesByUserID(userID int) ([]entity.Course, error) {
	courses, err := usecase.Repo.GetCoursesByUserID(userID)
	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (usecase CourseUseCase) GetCoursesStatus(userID int) (map[string]interface{}, error) {
	data := make(map[string]interface{})

	coursesInProgress, err := usecase.Repo.CoursesInProgress(userID)
	if err != nil {
		return nil, err
	}
	data["in_progress"] = coursesInProgress

	coursesDone, err := usecase.Repo.CoursesDone(userID)
	if err != nil {
		return nil, err
	}
	data["selesai"] = coursesDone

	return data, nil
}

func (usecase CourseUseCase) GetAllModules() ([]entity.Module, error) {
	modules, err := usecase.Repo.GetAllModules()
	return modules, err
}

func (usecase CourseUseCase) GetModule(id int) (entity.Module, error) {
	module, err := usecase.Repo.GetModule(id)
	return module, err
}
func (usecase CourseUseCase) GetAllCoursesWithSectionCount(courseId int) ([]entity.CourseWithSectionCount, error) {
	totalCourse, err := usecase.Repo.GetAllCoursesWithSectionCount(courseId)
	return totalCourse, err
}

func (usecase CourseUseCase) GetStudentsByCourseID(courseID int) ([]entity.User, error) {
	users, err := usecase.Repo.GetStudentsByCourseID(courseID)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (usecase CourseUseCase) GetAllCoursesWithSectionAndStudentCount(mentorId int) ([]entity.CourseWithSectionAndStudentCount, error) {
	courses, err := usecase.Repo.GetAllCoursesWithSectionAndStudentCount()
	if err != nil {
		return nil, err
	}

	return courses, nil
}
