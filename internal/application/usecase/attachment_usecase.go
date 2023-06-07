package usecase

import (
	"capston-lms/internal/adapters/repository"
	"capston-lms/internal/entity"
)

type AttachmentUseCase struct {
	Repo repository.AttachmentRepository
}

func (usecase AttachmentUseCase) GetAllAttachments(folderId int) ([]entity.Attachment, error) {
	Attachmentes, err := usecase.Repo.GetAllAttachments(folderId)
	return Attachmentes, err
}

func (usecase AttachmentUseCase) GetAttachment(id int) (entity.Attachment, error) {
	Attachment, err := usecase.Repo.GetAttachment(id)
	return Attachment, err
}

func (usecase AttachmentUseCase) CreateAttachment(Attachment entity.Attachment) error {
	err := usecase.Repo.CreateAttachment(Attachment)
	return err
}

func (usecase AttachmentUseCase) UpdateAttachment(id int, Attachment entity.Attachment) error {
	err := usecase.Repo.UpdateAttachment(id, Attachment)
	return err
}

func (usecase AttachmentUseCase) DeleteAttachment(id int) error {
	err := usecase.Repo.DeleteAttachment(id)
	return err
}
func (usecase AttachmentUseCase) FindAttachment(id int) error {
	err := usecase.Repo.FindAttachment(id)
	return err
}
