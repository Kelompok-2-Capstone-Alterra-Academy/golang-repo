package handler

import (
	"net/http"
	"strconv"

	"capston-lms/internal/application/service"
	"capston-lms/internal/application/usecase"
	"capston-lms/internal/entity"

	"github.com/labstack/echo/v4"
)

type AttachmentHandler struct {
	AttachmentUsecase usecase.AttachmentUseCase
}

func (handler AttachmentHandler) GetAllAttachments() echo.HandlerFunc {
	return func(e echo.Context) error {
		var categories []entity.Attachment
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     err.Error(),
			})
		}

		categories, err = handler.AttachmentUsecase.GetAllAttachments(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "success get all Attachments",
			"data":        categories,
		})
	}
}

func (handler AttachmentHandler) GetAttachment() echo.HandlerFunc {
	return func(e echo.Context) error {
		var Attachment entity.Attachment
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     err.Error(),
			})
		}

		Attachment, err = handler.AttachmentUsecase.GetAttachment(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "success get Attachment by id",
			"data":        Attachment,
		})
	}
}

func (handler AttachmentHandler) CreateAttachment() echo.HandlerFunc {
	return func(e echo.Context) error {
		var Attachment entity.Attachment
		if err := e.Bind(&Attachment); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     err.Error(),
			})
		}

		if Attachment.Type == "document" {
			file, err := e.FormFile("attachment_source")
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
			Attachment.AttachmentSource = result
		}

		err := handler.AttachmentUsecase.CreateAttachment(Attachment)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     "failed to create attachment",
			})
		}

		data := make(map[string]interface{})
		data["attachment"] = Attachment

		return e.JSON(
			http.StatusCreated, map[string]interface{}{
				"status code": http.StatusCreated,
				"message":     "success create new Attachment",
				"data":        data,
			})
	}
}

func (handler AttachmentHandler) UpdateAttachment() echo.HandlerFunc {
	var Attachment entity.Attachment

	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     err.Error(),
			})
		}

		err = handler.AttachmentUsecase.FindAttachment(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		if err := e.Bind(&Attachment); err != nil {
			return e.JSON(http.StatusNotFound, map[string]interface{}{
				"status code": http.StatusNotFound,
				"message":     err.Error(),
			})
		}

		err = handler.AttachmentUsecase.UpdateAttachment(id, Attachment)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "success update Attachment",
			"data":        Attachment,
		})
	}
}
func (handler AttachmentHandler) DeleteAttachment() echo.HandlerFunc {
	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     "input id is not number",
			})
		}

		err = handler.AttachmentUsecase.DeleteAttachment(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "Success Delete Attachment`",
		})
	}
}

func (handler AttachmentHandler) GetQuizAttachments(c echo.Context) error {
	attachments, err := handler.AttachmentUsecase.GetQuizAttachments()
	if err != nil {
		// Handle error from use case
		return c.String(http.StatusInternalServerError, "Failed to get quiz attachments")
	}

	return c.JSON(http.StatusOK, attachments)
}

func (handler AttachmentHandler) GetQuizAttachmentByID(c echo.Context) error {
	id := c.Param("id")

	// Convert the ID to an integer
	attachmentID, err := strconv.Atoi(id)
	if err != nil {
		// Handle the error
		return c.String(http.StatusBadRequest, "Invalid attachment ID")
	}

	attachment, err := handler.AttachmentUsecase.GetMateriAttachmentByID(attachmentID)
	if err != nil {
		// Handle the error
		return c.String(http.StatusInternalServerError, "Failed to get quiz attachment")
	}

	// Return the selected fields of the attachment as a response
	return c.JSON(http.StatusOK, map[string]interface{}{
		"attachment_name":   attachment.AttachmentName,
		"description":       attachment.Description,
		"attachment_source": attachment.AttachmentSource,
	})
}

func (handler AttachmentHandler) GetMateriAttachments(c echo.Context) error {
	attachments, err := handler.AttachmentUsecase.GetMateriAttachments()
	if err != nil {
		// Handle error from use case
		return c.String(http.StatusInternalServerError, "Failed to get materi attachments")
	}

	return c.JSON(http.StatusOK, attachments)
}

func (handler AttachmentHandler) GetMateriAttachmentByID(c echo.Context) error {
	id := c.Param("id")

	// Convert the ID to an integer
	attachmentID, err := strconv.Atoi(id)
	if err != nil {
		// Handle the error
		return c.String(http.StatusBadRequest, "Invalid attachment ID")
	}

	attachment, err := handler.AttachmentUsecase.GetMateriAttachmentByID(attachmentID)
	if err != nil {
		// Handle the error
		return c.String(http.StatusInternalServerError, "Failed to get materi attachment")
	}

	// Return the selected fields of the attachment as a response
	return c.JSON(http.StatusOK, map[string]interface{}{
		"attachment_name":   attachment.AttachmentName,
		"description":       attachment.Description,
		"attachment_source": attachment.AttachmentSource,
	})
}
