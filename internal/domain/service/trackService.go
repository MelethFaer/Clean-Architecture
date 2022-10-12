package service

import (
	"SandBox/internal/domain/entity"
	"SandBox/internal/domain/repository"
)

type TrackService struct {
	repo repository.Track
}

func NewTrackService(repo repository.Track) *TrackService {
	return &TrackService{repo: repo}
}

func (s *TrackService) GetOne(id string) (*entity.Track, error) {
	return s.repo.GetOne(id)
}

func (s *TrackService) GetUserPlaylist(id string) (*[]entity.Track, error) {
	return s.repo.GetByUser(id)
}

func (s *TrackService) GetAll() (*[]entity.Track, error) {
	return s.repo.GetAll()
}

func (s *TrackService) Create(user *entity.Track) (*entity.Track, error) {
	return s.repo.Create(user)
}

func (s *TrackService) Delete(id string) error {
	return s.repo.Delete(id)
}
