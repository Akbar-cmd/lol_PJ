package handlers

import (
	taskService2 "Poehali/internal/taskService"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	Service *taskService2.TaskService
}

// Нужна для создания структуры Handler на этапе инициализации приложения

func NewHandler(service *taskService2.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetTasksHandler(c echo.Context) error {
	tasks, err := h.Service.GetAllTasks()
	if err != nil {
		return c.JSON(http.StatusBadRequest, taskService2.Response{
			Status:  "error",
			Message: "Could not add the message",
		})
	}
	return c.JSON(http.StatusOK, tasks)
}

func (h *Handler) PostTaskHandler(c echo.Context) error {
	var task taskService2.Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, taskService2.Response{ //Декодируем JSON-тело запроса
			Status:  "Error",
			Message: "Could not add the message",
		})
	}
	createdTask, err := h.Service.CreateTask(task) //Создаем задачу через сервис
	if err != nil {
		return c.JSON(http.StatusBadRequest, taskService2.Response{
			Status:  "Error",
			Message: "Could not add the message",
		})
	}
	return c.JSON(http.StatusOK, createdTask)
}

func (h *Handler) PatchtaskHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, taskService2.Response{
			Status:  "Error",
			Message: "Bad ID",
		})
	}

	var updatedMessage taskService2.Task
	if err := c.Bind(&updatedMessage); err != nil {
		return c.JSON(http.StatusBadRequest, taskService2.Response{
			Status:  "Error",
			Message: "Could not update the message",
		})
	}
	// Обновляем задачу в базе данных
	task, err := h.Service.UpdateTaskByID(uint(id), updatedMessage)
	if err != nil {
		return c.JSON(http.StatusBadRequest, taskService2.Response{
			Status:  "Error",
			Message: "Could not update the message",
		})
	}

	// Возвращаем обновленную задачу
	return c.JSON(http.StatusOK, task)
}

func (h *Handler) DeleteHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, taskService2.Response{
			Status:  "Error",
			Message: "Bad ID",
		})
	}

	if err := h.Service.DeleteTaskByID(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, taskService2.Response{
			Status:  "Error",
			Message: "Could not delete the message",
		})
	}
	return c.NoContent(http.StatusNoContent)
}
