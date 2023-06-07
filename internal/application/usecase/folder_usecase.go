package usecase

import (
	"capston-lms/internal/adapters/repository"
	"capston-lms/internal/entity"
)

type FolderUseCase struct {
	Repo repository.FolderRepository
}

func (usecase FolderUseCase) GetAllFolders(mentorId int) ([]entity.Folder, error) {
	Folderes, err := usecase.Repo.GetAllFolders(mentorId)
	return Folderes, err
}

func (usecase FolderUseCase) GetFolder(id int) (entity.Folder, error) {
	Folder, err := usecase.Repo.GetFolder(id)
	return Folder, err
}

func (usecase FolderUseCase) CreateFolder(Folder entity.Folder) error {
	err := usecase.Repo.CreateFolder(Folder)
	return err
}

func (usecase FolderUseCase) UpdateFolder(id int, Folder entity.Folder) error {
	err := usecase.Repo.UpdateFolder(id, Folder)
	return err
}

func (usecase FolderUseCase) DeleteFolder(id int) error {
	err := usecase.Repo.DeleteFolder(id)
	return err
}
func (usecase FolderUseCase) FindFolder(id int) error {
	err := usecase.Repo.FindFolder(id)
	return err
}
