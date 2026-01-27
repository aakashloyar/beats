package service

import (
	"context"
	"github.com/aakashloyar/beats/track/internal/application/ports/in/track"
	"github.com/aakashloyar/beats/track/internal/application/ports/out"
	"github.com/aakashloyar/beats/track/internal/domain"
)

type ListTracksService struct {
	trackRepo out.TrackRepository
}

func NewListTracksService(trackRepo out.TrackRepository) in.ListTracksService {
	return &ListTracksService{
		trackRepo: trackRepo,
	}
}

func (s *ListTracksService) Execute(ctx context.Context, input in.ListTracksInput) ([]in.ListTrackOutput, error) {
	x, err := s.trackRepo.ListTracks(domain.TrackFilter{
		Title:    &input.Title,
		ArtistID: &input.ArtistID,
		AlbumID:  &input.AlbumID,
		Limit:    &input.Limit,
		Offset:   &input.Offset,
	})
	if err != nil {
		return nil, err
	}
	var tracks []in.ListTrackOutput
	for _, each := range x {
		track := in.ListTrackOutput{
			ID:            each.ID,
			ArtistID:      each.ArtistID,
			AlbumID:       each.AlbumID,
			CoverImageURL: each.CoverImageURL,
			DurationMS:    each.DurationMS,
			Language:      each.Language,
			ReleaseDate:   each.ReleaseDate,
			CreatedAt:     each.CreatedAt,
		}
		tracks = append(tracks, track)
	}

	return tracks, nil
}
