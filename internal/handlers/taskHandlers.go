package handlers

import (
	"Poehali/internal/taskService"
	"Poehali/internal/userService"
	"Poehali/internal/web/tasks"
	"context"
	"errors"
	"log"
)

type Handler struct {
	Service     *taskService.TaskService
	UserService *userService.UserService
}

// Нужна для создания структуры Handler на этапе инициализации приложения

func NewHandler(service *taskService.TaskService, userService *userService.UserService) *Handler {
	return &Handler{
		Service:     service,
		UserService: userService,
	}
}

func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	// Получение всех задач из сервиса
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := tasks.GetTasks200JSONResponse{}

	// Заполняем слайс response всеми задачами из БД
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}
	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	// Получаем userID из запроса
	userID := *request.Body.UserId

	// Создаем задачу только с UserID
	createdTask, err := h.Service.PostTask(userID, *request.Body.Task)
	if err != nil {
		return nil, err
	}

	// Формируем ответ
	response := tasks.PostTasks200JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task, // Может быть nil или пустой строкой
		IsDone: &createdTask.IsDone,
		UserId: &createdTask.UserID,
	}
	return response, nil
}

// Обновление задачи по ID
func (h *Handler) PatchTasksId(_ context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	// ID приходит из request.PathParams
	id := request.Id

	updatedTask := taskService.Task{}

	task, err := h.Service.UpdateTaskByID(uint(id), updatedTask)
	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     &task.ID,
		Task:   &task.Task,
		IsDone: &task.IsDone,
	}

	return response, nil
}

// Удаление задачи по ID
func (h *Handler) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	id := request.Id

	if err := h.Service.DeleteTaskByID(uint(id)); err != nil {
		return nil, err
	}

	return tasks.DeleteTasksId204Response{}, nil
}

func (h *Handler) GetUsersUserIdTasks(ctx context.Context, request tasks.GetUsersUserIdTasksRequestObject) (tasks.GetUsersUserIdTasksResponseObject, error) {
	userID := uint(request.UserId)

	if h.UserService == nil {
		log.Println("UserService is nil in GetUsersUserIdTasks")
		return nil, errors.New("internal server error")
	}

	tasksForUser, err := h.UserService.GetTasksByUserID(userID)
	if err != nil {
		return nil, err
	}

	response := tasks.GetUsersUserIdTasks200JSONResponse{}
	for _, tsk := range tasksForUser {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
			UserId: &tsk.UserID,
		}
		response = append(response, task)
	}

	return response, nil
}
