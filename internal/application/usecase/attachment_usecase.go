package usecase

import (
	"capston-lms/internal/adapters/repository"
	"capston-lms/internal/entity"
)

type AttachmentUseCase struct {
	Repo repository.AttachmentRepository
}

func (usecase AttachmentUseCase) GetAttachmentByCourseSectionIDAndAttachmentID(courseSectionID string, attachmentID string) (*entity.Attachment, error) {
	attachment, err := usecase.Repo.GetAttachmentByCourseSectionIDAndAttachmentID(courseSectionID, attachmentID)
	if err != nil {
		return nil, err
	}

	return attachment, nil
}
