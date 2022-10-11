package repository

import (
	"SandBox/internal/domain/entity"
	"gorm.io/gorm"
)

type TrackRepository struct {
	db *gorm.DB
}

func NewTrackRepository(db *gorm.DB) *TrackRepository {
	return &TrackRepository{db: db}
}

func(r *TrackRepository) GetOne(id uint) (*entity.Track, error) {
	track := &entity.Track{}
	if err := r.db.First(track, id).Error; err != nil {
		return nil, err
	}
	return track, nil
}

func(r *TrackRepository) GetAll() (*[]entity.Track, error) {
	tracks := &[]entity.Track{}
	if err := r.db.Find(tracks).Error; err != nil {
		return nil, err
	}
	return tracks, nil
}

func(r *TrackRepository) Create(track *entity.Track) (*entity.Track, error) {
	if err := r.db.Create(track).Error; err != nil {
		return nil, err
	}
	return track, nil
}

func(r *TrackRepository) Delete(id uint) error {
	if err := r.db.Delete(&entity.Track{}, id).Error; err != nil {
		return err
	}
	return nil
}
