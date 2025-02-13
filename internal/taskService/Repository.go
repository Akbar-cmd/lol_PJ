package taskService

import (
	"gorm.io/gorm"
)

type TaskRepository interface {
	//CreateTask - Передаем в функцию task типа Task из orm.go
	//Возвращаем созданный Task и ошибку
	CreateTask(task Task) (Task, error)
	//GetAllTasks - Возвращаем массив из всех задач в БД и ошибку
	GetAllTasks() ([]Task, error)
	//UpdateTaskByID - Передаем id и Task, возвращаем обновленный Task
	//и ошибку
	UpdateTaskByID(id uint, task Task) (Task, error)
	//DeleteTaskByID - Передаем id для удаления,возвращаем только ошибку
	DeleteTaskByID(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

// (r *taskRepository) привязывает данную функцию к нашему репозиторию
func (r *taskRepository) CreateTask(task Task) (Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTaskByID(id uint, task Task) (Task, error) {
	// обновляем задачу в БД
	if err := r.db.Model(&Task{}).Where("id = ?", id).Updates(task).Error; err != nil {
		return Task{}, err
	}
	// Получаем обновленную задачу из базы данных
	var updatedTask Task
	if err := r.db.First(&updatedTask, id).Error; err != nil {
		return Task{}, err
	}
	return updatedTask, nil
}

func (r *taskRepository) DeleteTaskByID(id uint) error {
	return r.db.Delete(&Task{}, id).Error
}
