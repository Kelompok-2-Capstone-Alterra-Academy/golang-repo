package handler

import (
	"net/http"
	"strconv"

	"capston-lms/internal/application/service"
	"capston-lms/internal/application/usecase"
	"capston-lms/internal/entity"

	"github.com/labstack/echo/v4"
)

type SubmissionHandler struct {
	SubmissionUseCase usecase.SubmissionUseCase
}

func (handler SubmissionHandler) GetAllSubmissions() echo.HandlerFunc {
	return func(e echo.Context) error {
		var Submissions []entity.Submission

		Submissions, err := handler.SubmissionUseCase.GetAllSubmissiones()
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "success get all Submissions",
			"data":        Submissions,
		})
	}
}

func (handler SubmissionHandler) GetSubmission() echo.HandlerFunc {
	return func(e echo.Context) error {
		var Submission entity.Submission
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     err.Error(),
			})
		}

		Submission, err = handler.SubmissionUseCase.GetSubmission(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "success get Submission by id",
			"data":        Submission,
		})
	}
}

func (handler SubmissionHandler) CreateSubmission() echo.HandlerFunc {
	return func(e echo.Context) error {

		var Submission entity.Submission
		if err := e.Bind(&Submission); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     err.Error(),
			})
		}
		StudentId, err := service.GetUserIDFromToken(e)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		file, err := e.FormFile("submission_source")
		if err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()
		result, err := service.UploadToS3(e, file.Filename, src)
		if err != nil {
			return err
		}

		Submission.SubmissionSource = result
		Submission.StudentId = StudentId

		err = handler.SubmissionUseCase.CreateSubmission(Submission)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     "failed to created Submission",
			})
		}
		return e.JSON(
			http.StatusCreated, map[string]interface{}{
				"status code": http.StatusCreated,
				"message":     "success create new Submission",
				"data":        Submission,
			})
	}
}

func (handler SubmissionHandler) UpdateSubmission() echo.HandlerFunc {
	var Submission entity.Submission
	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     err.Error(),
			})
		}

		err = handler.SubmissionUseCase.FindSubmission(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		if err := e.Bind(&Submission); err != nil {
			return e.JSON(http.StatusNotFound, map[string]interface{}{
				"status code": http.StatusNotFound,
				"message":     err.Error(),
			})
		}

		err = handler.SubmissionUseCase.UpdateSubmission(id, Submission)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "success update Submission",
			"data":        Submission,
		})
	}
}

func (handler SubmissionHandler) DeleteSubmission() echo.HandlerFunc {
	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     "input id is not number",
			})
		}

		err = handler.SubmissionUseCase.DeleteSubmission(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "Success Delete Submission`",
		})
	}
}
