package track 

import (
	"github.com/aakashloyar/beats/track/internal/application/ports/out"
	"github.com/aakashloyar/beats/track/internal/application/ports/in/track"
	"github.com/aakashloyar/beats/track/internal/domain"
	"context"
)

type ListTracksService struct {
	trackRepo out.TrackRepository
}

func NewListTracksService(trackRepo out.TrackRepository) in.ListTracksService {
	return &ListTracksService{
		trackRepo: trackRepo,
	}
}

func (s *ListTracksService) Execute(ctx context.Context,input in.ListTracksInput) ([]domain.Track,error) {
	x,err := s.trackRepo.ListTracks(input)
	if err != nil {
		return nil ,err 
	}
	return x, nil 
}

