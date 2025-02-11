package main

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Task   string `json:"task"` // Наш сервер будет ожидать json c полем task
	IsDone bool   `json:"is_done"`
}
