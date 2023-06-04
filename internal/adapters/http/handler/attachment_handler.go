package handler

import (
	"net/http"

	"capston-lms/internal/application/usecase"

	"github.com/labstack/echo/v4"
)

type AttachmentHandler struct {
	AttachmentUsecase usecase.AttachmentUseCase
}

func (handler AttachmentHandler) GetAttachmentByCourseSectionIDAndAttachmentID() echo.HandlerFunc {
	return func(c echo.Context) error {
		courseSectionID := c.Param("courseSectionID")
		attachmentID := c.Param("attachmentID")

		attachment, err := handler.AttachmentUsecase.GetAttachmentByCourseSectionIDAndAttachmentID(courseSectionID, attachmentID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "success get attachment by course section ID and attachment ID",
			"data":        attachment,
		})
	}
}
