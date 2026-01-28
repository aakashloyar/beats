package in

import (
	"context"
	"github.com/aakashloyar/beats/track/internal/domain"
	"time"
)

type CreateTrackInput struct {
	Title         string
	ArtistID      string
	AlbumID       *string
	CoverImageURL *string
	DurationMS    int64
	Language      domain.Language
	ReleaseDate   *time.Time
}

type CreateTrackOutput struct {
	TrackID string
}

type CreateTrackService interface {
	Execute(ctx context.Context, input CreateTrackInput) (CreateTrackOutput, error)
}
