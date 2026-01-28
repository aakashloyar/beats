package in

import (
	"context"
)

type CreateArtistInput struct {
	Name            string
	Bio             *string
	ProfileImageURL *string
}

type CreateArtistOutput struct {
	ArtistID string
}

type CreateArtistService interface {
	Execute(ctx context.Context, input CreateArtistInput) (CreateArtistOutput, error)
}
