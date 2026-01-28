package service

import (
	"context"
	"errors"
	"github.com/aakashloyar/beats/track/internal/application/ports/in/album"
	"github.com/aakashloyar/beats/track/internal/application/ports/out"
	"github.com/aakashloyar/beats/track/internal/domain"
)

type CreateAlbumService struct {
	albumRepo out.AlbumRepository
	idGen     domain.IDGenerator
	clock     domain.Clock
}

func NewCreateAlbumService(albumRepo out.AlbumRepository, idGen domain.IDGenerator, clock domain.Clock) in.CreateAlbumService {
	return &CreateAlbumService{
		albumRepo: albumRepo,
		idGen:     idGen,
		clock:     clock,
	}
}

func (s *CreateAlbumService) Execute(ctx context.Context, input in.CreateAlbumInput) (in.CreateAlbumOutput, error) {
	if input.Title == "" {
		return in.CreateAlbumOutput{}, errors.New("Name is required")
	}
	album := domain.Album{
		ID:            s.idGen.NewID(),
		Title:         input.Title,
		CoverImageURL: input.CoverImageURL,
		ReleaseDate:   input.ReleaseDate,
		CreatedAt:     s.clock.Now(),
	}
	if err := s.albumRepo.Save(album); err != nil {
		return in.CreateAlbumOutput{}, err
	}
	return in.CreateAlbumOutput{AlbumID: album.ID}, nil
}
