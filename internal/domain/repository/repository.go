package repository

import (
	"SandBox/internal/domain/entity"
	"gorm.io/gorm"
)

type Track interface {
	GetOne(id string) (*entity.Track, error)
	GetAll() (*[]entity.Track, error)
	GetByUser(id string) (*[]entity.Track, error)
	Create(track *entity.Track) (*entity.Track, error)
	Delete(id string) error
}

type User interface {
	GetOne(id string) (*entity.User, error)
	GetAll() (*[]entity.User, error)
	Create(user *entity.User) (*entity.User, error)
	Delete(id string) error
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
