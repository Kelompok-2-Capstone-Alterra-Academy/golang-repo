package usecase

import (
	"capston-lms/internal/adapters/repository"
	"capston-lms/internal/entity"
)

type SubmissionUseCase struct {
	Repo repository.SubmissionRepository
}

func (usecase SubmissionUseCase) GetAllSubmissiones() ([]entity.Submission, error) {
	Submissiones, err := usecase.Repo.GetAllSubmissions()
	return Submissiones, err
}

func (usecase SubmissionUseCase) GetSubmission(id int) (entity.Submission, error) {
	Submission, err := usecase.Repo.GetSubmission(id)
	return Submission, err
}

func (usecase SubmissionUseCase) CreateSubmission(Submission entity.Submission) error {
	err := usecase.Repo.CreateSubmission(Submission)
	return err
}

func (usecase SubmissionUseCase) UpdateSubmission(id int, Submission entity.Submission) error {
	err := usecase.Repo.UpdateSubmission(id, Submission)
	return err
}

func (usecase SubmissionUseCase) DeleteSubmission(id int) error {
	err := usecase.Repo.DeleteSubmission(id)
	return err
}
func (usecase SubmissionUseCase) FindSubmission(id int) error {
	err := usecase.Repo.FindSubmission(id)
	return err
}
