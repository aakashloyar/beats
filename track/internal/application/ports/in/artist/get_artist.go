package in

import (
	"context"
	"time"
)

type GetArtistOutput struct {
	ID              string
	Name            string
	Bio             *string
	ProfileImageURL *string
	CreatedAt       time.Time
}

type GetArtistInput struct {
	ArtistID string
}

type GetArtistService interface {
	Execute(ctx context.Context, input GetArtistInput) (GetArtistOutput, error)
}
