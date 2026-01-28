package service

import (
	"context"
	"github.com/aakashloyar/beats/track/internal/application/ports/in/album"
	"github.com/aakashloyar/beats/track/internal/application/ports/out"
)

type ListAlbumsService struct {
	albumRepo out.AlbumRepository
}

func NewListAlbumsService(albumRepo out.AlbumRepository) in.ListAlbumsService {
	return &ListAlbumsService{
		albumRepo: albumRepo,
	}
}

func (s *ListAlbumsService) Execute(ctx context.Context, input in.ListAlbumsInput) ([]in.ListAlbumsOutput, error) {
	x, err := s.albumRepo.ListAlbums(input.Title)
	if err != nil {
		return nil, err
	}
	var albums []in.ListAlbumsOutput
	for _, each := range x {
		album := in.ListAlbumsOutput{
			ID:            each.ID,
			Title:         each.Title,
			CoverImageURL: each.CoverImageURL,
			ReleaseDate:   each.ReleaseDate,
			CreatedAt:     each.CreatedAt,
		}
		albums = append(albums, album)
	}

	return albums, nil
}
