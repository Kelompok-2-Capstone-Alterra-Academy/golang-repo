package handler

import (
	"capston-lms/internal/application/service"
	"capston-lms/internal/application/usecase"
	"capston-lms/internal/entity"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TaskSubmissionHandler struct {
	TaskUseCase usecase.TaskUseCase
}

func submitTask(c echo.Context) error {
	// Get the student's ID from the request header or session
	studentID := c.Request().Header.Get("Student-ID")

	// Bind the request body to the TaskSubmission struct
	taskSubmission := new(entity.TaskSubmission)
	if err := c.Bind(taskSubmission); err != nil {
		return c.String(http.StatusBadRequest, "Failed to parse the request body")
	}

	// Upload the file to AWS S3
	if err := service.UploadToS3(c echo.Context, filename string, src multipart.File); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to upload the file to AWS S3")
}

	// Save the task submission in the database or perform any other necessary operations
	// Here, you can store the student ID, file URL from S3, and notes in your database

	// Return a success response
	response := fmt.Sprintf("Task submitted successfully. Student ID: %s, Notes: %s", studentID, taskSubmission.Notes)
	return c.String(http.StatusOK, response)
}