package repository

import (
	"capston-lms/internal/entity"

	"gorm.io/gorm"
)

type FolderRepository struct {
	DB *gorm.DB
}

func (repo FolderRepository) GetAllFolders(mentorId int) ([]entity.Folder, error) {
	var Folders []entity.Folder
	result := repo.DB.Where("mentor_id = ?", mentorId).Find(&Folders)
	return Folders, result.Error
}

func (repo FolderRepository) GetFolder(id int) (entity.Folder, error) {
	var Folders entity.Folder
	result := repo.DB.First(&Folders, id)
	return Folders, result.Error
}

func (repo FolderRepository) CreateFolder(Folder *entity.Folder) error {
	result := repo.DB.Create(&Folder)
	return result.Error
}

func (repo FolderRepository) UpdateFolder(id int, Folder entity.Folder) error {
	result := repo.DB.Model(&Folder).Where("id = ?", id).Updates(&Folder)
	return result.Error
}

func (repo FolderRepository) DeleteFolder(id int) error {
	result := repo.DB.Delete(&entity.Folder{}, id)
	return result.Error
}

func (repo FolderRepository) FindFolder(id int) error {
	result := repo.DB.First(&entity.Folder{}, id)
	return result.Error
}
