package handler

import (
	"net/http"
	"strconv"

	"capston-lms/internal/application/usecase"
	"capston-lms/internal/entity"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	TaskUseCase usecase.TaskUseCase
}

func (handler TaskHandler) GetAllTasks() echo.HandlerFunc {
	return func(e echo.Context) error {
		var Tasks []entity.Task

		Tasks, err := handler.TaskUseCase.GetAllTasks()
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "success get all Tasks",
			"data":        Tasks,
		})
	}
}

func (handler TaskHandler) GetTask() echo.HandlerFunc {
	return func(e echo.Context) error {
		var Task entity.Task
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     err.Error(),
			})
		}

		Task, err = handler.TaskUseCase.GetTask(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "success get Task by id",
			"data":        Task,
		})
	}
}

func (handler TaskHandler) CreateTask() echo.HandlerFunc {
	return func(e echo.Context) error {
		var Task entity.Task
		if err := e.Bind(&Task); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     err.Error(),
			})
		}

		err := handler.TaskUseCase.CreateTask(Task)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     "failed to created Task",
			})
		}
		return e.JSON(
			http.StatusCreated, map[string]interface{}{
				"status code": http.StatusCreated,
				"message":     "success create new Task",
				"data":        Task,
			})
	}
}
func (handler TaskHandler) UpdateTask() echo.HandlerFunc {
	var Task entity.Task

	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     err.Error(),
			})
		}

		err = handler.TaskUseCase.FindTask(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		if err := e.Bind(&Task); err != nil {
			return e.JSON(http.StatusNotFound, map[string]interface{}{
				"status code": http.StatusNotFound,
				"message":     err.Error(),
			})
		}

		err = handler.TaskUseCase.UpdateTask(id, Task)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "success update Task",
			"data":        Task,
		})
	}
}

func (handler TaskHandler) DeleteTask() echo.HandlerFunc {
	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message":     "input id is not number",
			})
		}

		err = handler.TaskUseCase.DeleteTask(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message":     err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message":     "Success Delete Task`",
		})
	}
}
