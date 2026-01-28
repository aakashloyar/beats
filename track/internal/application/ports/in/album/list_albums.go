package in

import (
	"context"
	"time"
)

type ListAlbumsInput struct {
	Title    string
}

type ListAlbumsOutput struct {
	ID            string
	Title         string
	CoverImageURL *string
	ReleaseDate   *time.Time
	CreatedAt     time.Time
}


type ListAlbumsService interface {
	Execute(ctx context.Context, input ListAlbumsInput) ([]ListAlbumsOutput, error)
}
