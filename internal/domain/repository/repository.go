package repository

import (
	"SandBox/internal/domain/entity"
	"gorm.io/gorm"
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

type Repository struct {
	Track
	User
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Track: NewTrackRepository(db),
		User:  NewUserRepository(db),
	}
}
