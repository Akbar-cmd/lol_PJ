package main

import (
	"Poehali/internal/database"
	"Poehali/internal/handlers"
	"Poehali/internal/taskService"
	userService2 "Poehali/internal/userService"
	"Poehali/internal/web/tasks"
	"Poehali/internal/web/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

// Основные методы GORM - Create, Find, Update, Delete (CRUD)

func main() {

	database.InitDB()

	tasksRepo := taskService.NewTaskRepository(database.DB)
	tasksService := taskService.NewService(tasksRepo)

	usersRepo := userService2.NewUserRepository(database.DB)
	usersService := userService2.NewUserService(usersRepo)

	tasksHandler := handlers.NewHandler(tasksService, usersService)
	usersHandler := handlers.NewUserHandler(usersService)

	// Инициализируем echo
	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	tasksStrictHandler := tasks.NewStrictHandler(tasksHandler, nil) // тут будет ошибка
	tasks.RegisterHandlers(e, tasksStrictHandler)

	usersStrictHandler := users.NewStrictHandler(usersHandler, nil)
	users.RegisterHandlers(e, usersStrictHandler)

	if err := e.Start(":8082"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
