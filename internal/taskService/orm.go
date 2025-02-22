package taskService

import "time"

type Task struct {
	ID        uint      `json:"id"`
	Task      string    `json:"task"`
	IsDone    bool      `json:"is_done"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`  //
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`  //`
	DeletedAt time.Time `gorm:"index" json:"deleted_at,omitempty"` //`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
