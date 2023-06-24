package handler

import (
	"net/http"
	"strconv"

	"capston-lms/internal/application/service"
	"capston-lms/internal/application/usecase"
	"capston-lms/internal/entity"

	"github.com/labstack/echo/v4"
)

type FolderHandler struct {
	FolderUsecase usecase.FolderUseCase
}

func (handler FolderHandler) GetAllFolders() echo.HandlerFunc {
	return func(e echo.Context) error {
		var folders []entity.Folder
		MentorId, err := service.GetUserIDFromToken(e)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		folders, err = handler.FolderUsecase.GetAllFolders(MentorId)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}
		data := make(map[string]interface{})
		data["folders"] = folders

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "success get all Folders",
			"data":        data,
		})
	}
}

func (handler FolderHandler) GetFolder() echo.HandlerFunc {
	return func(e echo.Context) error {
		var Folder entity.Folder
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     err.Error(),
			})
		}

		Folder, err = handler.FolderUsecase.GetFolder(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "success get Folder by id",
			"data":        Folder,
		})
	}
}

func (handler FolderHandler) CreateFolder() echo.HandlerFunc {
	return func(e echo.Context) error {
		var folder entity.Folder
		if err := e.Bind(&folder); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     err.Error(),
			})
		}

		MentorId, err := service.GetUserIDFromToken(e)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		folder.MentorId = MentorId
		err = handler.FolderUsecase.CreateFolder(&folder)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     "failed to created user",
			})
		}

		data := make(map[string]interface{})
		data["folder"] = folder

		return e.JSON(
			http.StatusCreated, map[string]interface{}{
				"status code": http.StatusCreated,
				"message":     "success create new Folder",
				"data":        data,
			})
	}
}
func (handler FolderHandler) UpdateFolder() echo.HandlerFunc {
	var Folder entity.Folder

	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     err.Error(),
			})
		}

		err = handler.FolderUsecase.FindFolder(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		if err := e.Bind(&Folder); err != nil {
			return e.JSON(http.StatusNotFound, map[string]interface{}{
				"status code": http.StatusNotFound,
				"message":     err.Error(),
			})
		}

		err = handler.FolderUsecase.UpdateFolder(id, Folder)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "success update Folder",
			"data":        Folder,
		})
	}
}
func (handler FolderHandler) DeleteFolder() echo.HandlerFunc {
	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     "input id is not number",
			})
		}

		err = handler.FolderUsecase.DeleteFolder(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "Success Delete Folder`",
		})
	}
}
