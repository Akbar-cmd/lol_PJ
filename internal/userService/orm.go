package userService

import (
	"Poehali/internal/taskService"
	"time"
)

type User struct {
	ID        uint               `json:"id"`
	Email     string             `json:"email"`
	Password  string             `json:"password"`
	CreatedAt time.Time          `gorm:"autoCreateTime" json:"created_at"`  //
	UpdatedAt time.Time          `gorm:"autoUpdateTime" json:"updated_at"`  //`
	DeletedAt time.Time          `gorm:"index" json:"deleted_at,omitempty"` //`
	Task      []taskService.Task `json:"task"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
