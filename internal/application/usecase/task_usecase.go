package usecase

import (
	"capston-lms/internal/adapters/repository"
	"capston-lms/internal/entity"
)

type TaskUseCase struct {
	Repo repository.TaskRepository
}

func (usecase TaskUseCase) GetAllTasks() ([]entity.Task, error) {
	Taskes, err := usecase.Repo.GetAllTasks()
	return Taskes, err
}

func (usecase TaskUseCase) GetTask(id int) (entity.Task, error) {
	Task, err := usecase.Repo.GetTask(id)
	return Task, err
}

func (usecase TaskUseCase) CreateTask(Task entity.Task) error {
	err := usecase.Repo.CreateTask(Task)
	return err
}

func (usecase TaskUseCase) UpdateTask(id int, Task entity.Task) error {
	err := usecase.Repo.UpdateTask(id, Task)
	return err
}

func (usecase TaskUseCase) DeleteTask(id int) error {
	err := usecase.Repo.DeleteTask(id)
	return err
}

func (usecase TaskUseCase) FindTask(id int) error {
	err := usecase.Repo.FindTask(id)
	return err
}
