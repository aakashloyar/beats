package track

import (
	"encoding/json"
	"net/http"
	"github.com/aakashloyar/beats/track/internal/application/ports/in"
	"github.com/aakashloyar/beats/track/internal/domain"
	"time"
)

type CreateTrackRequest struct {
	Title         string          `json:"title"`
	ArtistID      string          `json:"artist_id"`
	AlbumID       *string         `json:"album_id,omitempty"`
	CoverImageURL *string         `json:"cover_image_url,omitempty"`
	DurationMS    int64           `json:"duration_ms"`
	Language      domain.Language `json:"language"`
	ReleaseDate   *time.Time      `json:"release_date,omitempty"`
}

type CreateTrackResponse struct {
	TrackID string `json:"track_id"`
}

type GetTrackResponse struct {
	ID            string           `json:"id"`
	Title         string           `json:"title"`
	ArtistID      string           `json:"artist_id"`
	AlbumID       *string          `json:"album_id,omitempty"`
	CoverImageURL *string          `json:"cover_image_url,omitempty"`
	DurationMS    int64            `json:"duration_ms"`
	Language      domain.Language  `json:"language"`
	ReleaseDate   *time.Time       `json:"release_data,omitempty"`
	CreatedAt     time.Time        `json:"created_at"`
}

type Handler struct {
	createTrackService in.CreateTrackService 
	getTrackService    in.GetTrackService 
}

func NewHandler(createTrackService in.CreateTrackService,getTrackService in.GetTrackService) *Handler {
	return &Handler{
		createTrackService: createTrackService,
		getTrackService:    getTrackService,
	}
}


func (h *Handler) CreateTrack(w http.ResponseWriter, r *http.Request) {
	var req CreateTrackRequest 
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
	}
	input := &in.CreateTrackInput{
		Title:         req.Title,
		ArtistID:      req.ArtistID,
		AlbumID:       req.AlbumID,
		CoverImageURL: req.CoverImageURL,
		DurationMS:    req.DurationMS,
		Language:      req.Language,
		ReleaseDate:   req.ReleaseDate,
	}
	out, err := h.createTrackService.Execute(r.Context(),input)

	if err != nil {
		http.Error(w, err.Error(),http.StatusInternalServerError)
		return 
	}
	resp := CreateTrackResponse{
		TrackID: out.TrackID,
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) GetTrackById(w http.ResponseWriter, r *http.Request, id string) {
	out, err := h.getTrackService.Execute(r.Context(),id)
	if err != nil {
		http.Error(w, err.Error(),http.StatusInternalServerError)
		return 
	}
	resp := GetTrackResponse{
		ID: out.ID,
		Title: out.Title,
		ArtistID: out.ArtistID,
		AlbumID: out.AlbumID,
		CoverImageURL: out.CoverImageURL,
		DurationMS: out.DurationMS,
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
