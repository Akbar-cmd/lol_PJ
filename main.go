package main

import (
	"Poehali/internal/database"
	"Poehali/internal/handlers"
	"Poehali/internal/taskService"
	"github.com/labstack/echo/v4"
)

//Основные методы GORM - Create, Find, Update, Delete (CRUD)

func main() {
	//Вызываем InitDB() из файла db.go
	database.InitDB()
	//Автоматическая миграция модели Message
	database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)

	handler := handlers.NewHandler(service)

	e := echo.New()

	e.GET("/task", handler.GetTasksHandler)
	e.POST("/task", handler.PostTaskHandler)
	e.PATCH("/task/:id", handler.PatchtaskHandler)
	e.DELETE("/task/:id", handler.DeleteHandler)
	e.Start(":8080")
}
