package service

import (
	"context"

	"github.com/aakashloyar/beats/track/internal/application/ports/in/artist"
	"github.com/aakashloyar/beats/track/internal/application/ports/out"
)

type GetAlbumService struct {
	albumRepo out.AlbumRepository
}

func NewGetAlbumService(albumRepo out.AlbumRepository) in.GetArtistService {
	return &GetAlbumService{
		albumRepo: albumRepo,
	}
}

func (s *GetAlbumService) Execute(ctx context.Context, input in.GetAlbumInput) (in.GetAlbumOutput, error) {
	x, err := s.albumRepo.FindByID(input.AlbumID)
	if err != nil {
		return in.GetAlbumOutput{}, nil
	}
	album := in.GetAlbumOutput{
		ID:              x.ID,
		Title:           x.Title,
		CoverImageURL:   x.CoverImageURL,
		ProfileImageURL: x.ReleaseDate,
		CreatedAt:       x.CreatedAt,
	}
	return album, nil
}
