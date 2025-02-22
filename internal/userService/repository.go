package userService

import (
	"Poehali/internal/taskService"
	"errors"
	"gorm.io/gorm"
	"log"
)

type UserRepository interface {
	CreateUser(user User) (User, error)
	GetUser() ([]User, error)
	UpdateUserByID(id uint, user User) (User, error)
	DeleteUserByID(id uint) error
	GetTasksByUserID(userID uint) ([]taskService.Task, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository { return &userRepository{db: db} }

func (r *userRepository) CreateUser(user User) (User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (r *userRepository) GetUser() ([]User, error) {
	var users []User               //создаем слайс для хранения пользователей
	err := r.db.Find(&users).Error //запрашиваем всех пользователей из БД
	return users, err              //возвращаем список пользователей и ошибку(если есть)
}

func (r *userRepository) UpdateUserByID(id uint, user User) (User, error) {
	if err := r.db.Model(&User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		return User{}, err
	}
	var updatedUser User
	if err := r.db.First(&updatedUser, id).Error; err != nil {
		return User{}, err
	}
	return updatedUser, nil

}

func (r *userRepository) DeleteUserByID(id uint) error { return r.db.Delete(&User{}, id).Error }

func (r *userRepository) GetTasksByUserID(userID uint) ([]taskService.Task, error) {

	if r == nil {
		log.Println("userRepository is nil")
		return nil, errors.New("internal server error")
	}
	if r.db == nil {
		log.Println("Database connection is nil")
		return nil, errors.New("database not initialized")
	}

	var tasks []taskService.Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	if err != nil {
		log.Printf("Ошибка получения задач для userID %d: %v", userID, err)
		return nil, err
	}
	return tasks, nil
}
