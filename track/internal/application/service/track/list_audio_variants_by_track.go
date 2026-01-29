package service

import (
	"context"

	"github.com/aakashloyar/beats/track/internal/application/ports/in/track"
	"github.com/aakashloyar/beats/track/internal/application/ports/out"
)

type ListAudioVariantsByTrackService struct {
	repo out.TrackRepository
}

func NewListAudioVariantsByTrackService(repo out.TrackRepository) in.ListAudioVariantsByTrackService {
	return &ListAudioVariantsByTrackService{
		repo: repo,
	}
}

func (s *ListAudioVariantsByTrackService) Execute(ctx context.Context, input in.ListAudioVariantsByTrackInput) ([]in.ListAudioVariantsByTrackOutput, error) {

	x, err := s.repo.ListAudioVariantsByTrack(input.TrackID)
	if err != nil {
		return nil, err
	}
	var audiovariants []in.ListAudioVariantsByTrackOutput
	for _, each := range x {
		audiovariant := in.ListAudioVariantsByTrackOutput{
			ID:           each.ID,
			TrackID:      each.TrackID,
			Codec:        each.Codec,
			BitrateKbps:  each.BitrateKbps,
			SampleRateHz: each.SampleRateHz,
			Channels:     each.Channels,
			DurationMs:   each.DurationMs,
			FileURL:      each.FileURL,
			CreatedAt:    each.CreatedAt,
		}
		audiovariants = append(audiovariants, audiovariant)
	}
	return audiovariants, nil
}
