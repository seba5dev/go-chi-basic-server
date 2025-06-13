package service

import (
	"go-chi-basic-server/internal/models"
	"go-chi-basic-server/internal/repository"
)

// SongsService defines the interface for interacting with songs data.
type SongsService interface {
	// GetAll returns all songs
	GetAll() (map[int]models.Song, error)
}

// NewSongsDefault creates a new instance of SongsDefault with the provided repository.
func NewSongsDefault(rp repository.SongsRepository) *SongsDefault {
	return &SongsDefault{
		rp: rp,
	}
}

// SongsDefault is a struct that implements the SongsService interface.
type SongsDefault struct {
	// rp is the repository for songs
	rp repository.SongsRepository
}

// NewSongsDefault creates a new instance of SongsDefault with the provided repository.
func (s *SongsDefault) GetAll() (map[int]models.Song, error) {
	// get all songs from the repository
	return s.rp.GetAll()
}
