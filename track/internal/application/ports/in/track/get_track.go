package in

import (
	"context"
	"github.com/aakashloyar/beats/track/internal/domain"
	"time"
)

type GetTrackOutput struct {
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

type GetTrackInput struct {
	TrackID string
}
type GetTrackService interface {
	Execute(ctx context.Context, input GetTrackInput) (GetTrackOutput, error)
}
