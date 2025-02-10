package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

var task string

type requestBody struct {
	Message string `json:"message"`
}

// принимает данные от клиента
func PostHandler(c echo.Context) error {
	var r requestBody
	if err := c.Bind(&r); err != nil {
		return c.JSON(http.StatusBadRequest, "Ошибка")
	}
	task = r.Message
	return c.JSON(http.StatusOK, "Нармална")
}

// возвращает данные клиенту
func GetHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello,"+task)
}

func main() {
	e := echo.New()

	e.GET("/task", GetHandler)
	e.POST("/task", PostHandler)

	e.Start(":8080")
}
