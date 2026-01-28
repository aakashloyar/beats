package in

import (
	"context"
	"time"
)

type CreateAlbumInput struct {
	Title         string
	CoverImageURL *string
	ReleaseDate   *time.Time
}

type CreateAlbumOutput struct {
	AlbumID string
}

type CreateAlbumService interface {
	Execute(ctx context.Context, input CreateAlbumInput) (CreateAlbumOutput, error)
}
