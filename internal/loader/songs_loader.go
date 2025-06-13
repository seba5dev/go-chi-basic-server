package loader

import (
	"encoding/json"
	"go-chi-basic-server/internal/models"
	"os"
)

// SongsLoader defines the interface for loading songs data from a file.
type SongsLoader interface {
	// Load loads the songs from the file and returns a map of songs
	Load() (map[int]models.Song, error)
}

// SongsJSONFile is a struct that implements the SongsLoader interface for loading songs from a JSON file.
type SongsJSONFile struct {
	// FilePath is the path to the JSON file containing the songs data
	FilePath string
}

// Load loads the songs from the JSON file and returns a map of songs.
func NewSongsJSONFile(filePath string) *SongsJSONFile {
	// Create a new instance of SongsJSONFile with the provided file path
	return &SongsJSONFile{
		FilePath: filePath,
	}
}

// Load reads the JSON file and unmarshals it into a map of songs.
func (s *SongsJSONFile) Load() (map[int]models.Song, error) {
	// reads the JSON file and unmarshals it into a map of songs
	data, err := os.ReadFile(s.FilePath)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into a slice of SongsDoc
	var songsJson []models.SongsDoc
	err = json.Unmarshal(data, &songsJson)
	if err != nil {
		return nil, err
	}

	// Convert the slice of SongsDoc to a map of models.Song
	songsMap := make(map[int]models.Song)
	// Iterate over the slice and populate the map
	for _, song := range songsJson {
		songsMap[song.ID] = models.Song{
			ID:        song.ID,
			Title:     song.Title,
			Artist_id: song.Artist,
			Album:     song.Album,
			Year:      song.Year,
			Genre:     song.Genre,
		}
	}
	// Return the map of songs
	return songsMap, nil
}
