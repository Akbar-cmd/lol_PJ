package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

//requestBody

type Tasks struct {
	Message string `json:"task"`
	IsDone  bool   `json:"is_done"`
	ID      int    `json:"id"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

//Основные методы GORM - Create, Find, Update, Delete (CRUD)

// Возвращает данные клиенту
func GetHandler(c echo.Context) error {
	var task []Tasks
	if err := DB.Find(&task).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Could not add the message",
		})
	}
	return c.JSON(http.StatusOK, task)
}

// Принимает данные от клиента
func PostHandler(c echo.Context) error {
	var task Tasks
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not add the message",
		})
	}
	if err := DB.Create(&task).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not add the message",
		})
	}
	return c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Message was successfully created",
	})
}

// Обновляет данные по ID
func PatchHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Bad ID",
		})
	}

	var updatedMessage Tasks
	if err := c.Bind(&updatedMessage); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not update the message",
		})
	}

	if err := DB.Model(&Tasks{}).Where("id = ?", id).Update("task", updatedMessage.Message).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not update the message",
		})
	}

	return c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Message was updated",
	})
}

// Удаляет данные по ID
func DeleteHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Bad ID",
		})
	}

	if err := DB.Delete(&Tasks{}, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not delete the message",
		})
	}
	return c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Message was deleted",
	})
}

func main() {
	//Вызываем InitDB() из файла db.go
	initDB()
	//Автоматическая миграция модели Message
	DB.AutoMigrate(&Tasks{})

	e := echo.New()

	e.GET("/task", GetHandler)
	e.POST("/task", PostHandler)
	e.PATCH("/task/:id", PatchHandler)
	e.DELETE("/task/:id", DeleteHandler)
	e.Start(":8080")
}
