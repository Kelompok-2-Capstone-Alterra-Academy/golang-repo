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

func (usecase AttachmentUseCase) GetAllQuiz() ([]entity.Attachment, error) {
	Attachmentes, err := usecase.Repo.GetQuiz()
	return Attachmentes, err
}

func (usecase AttachmentUseCase) GetAttachment(id int) (entity.Attachment, error) {
	attachment, err := usecase.Repo.GetAttachment(id)
	return attachment, err
}

func (usecase AttachmentUseCase) CreateAttachment(Attachment *entity.Attachment) error {
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

func (usecase AttachmentUseCase) GetVideoAttachments() ([]entity.Attachment, error) {
	return usecase.Repo.GetVideoAttachments()
}

func (usecase AttachmentUseCase) GetVideoAttachmentByID(id int) (entity.Attachment, error) {
	attachment, err := usecase.Repo.GetVideoAttachmentByID(id)
	if err != nil {
		return entity.Attachment{}, err
	}
	return attachment, nil
}

func (usecase AttachmentUseCase) GetQuizAttachments() ([]entity.Attachment, error) {
	return usecase.Repo.GetQuizAttachments()
}

func (usecase AttachmentUseCase) GetQuizAttachmentByID(id int) (entity.Attachment, error) {
	attachment, err := usecase.Repo.GetQuizAttachmentByID(id)
	if err != nil {
		return entity.Attachment{}, err
	}
	return attachment, nil
}

func (usecase AttachmentUseCase) GetMateriAttachments() ([]entity.Attachment, error) {
	return usecase.Repo.GetMateriAttachments()
}

func (usecase AttachmentUseCase) GetMateriAttachmentByID(id int) (entity.Attachment, error) {
	attachment, err := usecase.Repo.GetMateriAttachmentByID(id)
	if err != nil {
		return entity.Attachment{}, err
	}
	return attachment, nil
}
