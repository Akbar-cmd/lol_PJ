package userService

import (
	"Poehali/internal/taskService"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService { return &UserService{repo: repo} }

func (s *UserService) CreateUser(user User) (User, error) { return s.repo.CreateUser(user) }

func (s *UserService) GetUser() ([]User, error) { return s.repo.GetUser() }

func (s *UserService) UpdateUserByID(id uint, user User) (User, error) {
	return s.repo.UpdateUserByID(id, user)
}

func (s *UserService) DeleteUserByID(id uint) error { return s.repo.DeleteUserByID(id) }

func (s *UserService) GetTasksByUserID(userID uint) ([]taskService.Task, error) {
	return s.repo.GetTasksByUserID(userID)
}
