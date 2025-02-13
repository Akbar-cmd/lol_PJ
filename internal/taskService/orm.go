package taskService

type Task struct {
	Message string `json:"task"` // Наш сервер будет ожидать json c полем task
	IsDone  bool   `json:"is_done"`
	ID      int    `json:"id"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
