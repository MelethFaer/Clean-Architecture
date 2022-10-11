package service

import (
	"SandBox/internal/domain/entity"
	"SandBox/internal/domain/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetOne(id uint) (*entity.User, error) {
	return s.repo.GetOne(id)
}

func (s *UserService) GetAll() (*[]entity.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) Create(user *entity.User) (*entity.User, error) {
	return s.repo.Create(user)
}

func (s *UserService) Delete(id uint) error {
	return s.repo.Delete(id)
}
