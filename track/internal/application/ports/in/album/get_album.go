package in

import (
	"context"
	"time"
)

type GetAlbumOutput struct {
	ID            string
	Title         string
	CoverImageURL *string
	ReleaseDate   *time.Time
	CreatedAt     time.Time
}

type GetAlbumInput struct {
	AlbumID string
}

type GetAlbumService interface {
	Execute(ctx context.Context, input GetAlbumInput) (GetAlbumOutput, error)
}
