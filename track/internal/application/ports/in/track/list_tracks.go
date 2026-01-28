package in

import (
	"context"
	"github.com/aakashloyar/beats/track/internal/domain"
	"time"
)

type ListTracksInput struct {
	Title    string
	ArtistID string
	AlbumID  string
	Limit    string
	Offset   string
}

type ListTracksOutput struct {
	ID            string
	Title         string
	ArtistID      string
	AlbumID       *string
	CoverImageURL *string
	DurationMS    int64
	Language      domain.Language
	ReleaseDate   *time.Time
	CreatedAt     time.Time
}

type ListTracksService interface {
	Execute(ctx context.Context, input ListTracksInput) ([]ListTracksOutput, error)
}
