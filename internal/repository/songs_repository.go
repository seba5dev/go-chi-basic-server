package repository

import "go-chi-basic-server/internal/models"

// SongsRepository defines the interface for interacting with songs data.
type SongsRepository interface {
	// GetAll returns all songs
	GetAll() (map[int]models.Song, error)
}

// SongsMap is an in-memory implementation of the SongsRepository interface.
type SongsMap struct {
	db map[int]models.Song
}

// NewSongsMap creates a new instance of SongsMap with the provided database.
func NewSongsMap(db map[int]models.Song) SongsRepository {
	defaultDb := make(map[int]models.Song)
	if db != nil {
		defaultDb = db
	}
	return &SongsMap{
		db: defaultDb,
	}
}

// GetAll retrieves all songs from the in-memory database.
func (s *SongsMap) GetAll() (map[int]models.Song, error) {
	v := make(map[int]models.Song)

	// copy the values from the map
	for k, v1 := range s.db {
		v[k] = v1
	}
	return v, nil
}
