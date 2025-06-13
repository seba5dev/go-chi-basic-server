package handler

import (
	"go-chi-basic-server/internal/models"
	"go-chi-basic-server/internal/service"
	"go-chi-basic-server/pkg/response"
	"net/http"
)

// SongsHandler defines the interface for handling songs requests.
func NewSongDefault(sv service.SongsService) *SongsDefault {
	return &SongsDefault{
		sv: sv,
	}
}

// SongsDefault is a struct that implements the SongsHandler interface.
type SongsDefault struct {
	// sv is the service for songs
	sv service.SongsService
}

func (s *SongsDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request

		// process
		songs, err := s.sv.GetAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]models.Song)
		for k, v := range songs {
			data[k] = models.Song{
				ID:        v.ID,
				Title:     v.Title,
				Artist_id: v.Artist_id,
				Album:     v.Album,
				Year:      v.Year,
				Genre:     v.Genre,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}
