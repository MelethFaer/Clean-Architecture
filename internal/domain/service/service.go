package service

import (
	"SandBox/internal/domain/entity"
	"SandBox/internal/domain/repository"
)

type Track interface {
	GetOne(id uint) (*entity.Track, error)
	GetAll() (*[]entity.Track, error)
	Create(track *entity.Track) (*entity.Track, error)
	Delete(id uint) error
}

type User interface {
	GetOne(id uint) (*entity.User, error)
	GetAll() (*[]entity.User, error)
	Create(user *entity.User) (*entity.User, error)
	Delete(id uint) error
}

type Service struct {
	Track
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Track: NewTrackService(repos.Track),
		User:  NewUserService(repos.User),
	}
}
