package in

import (
	"context"
	"github.com/aakashloyar/beats/track/internal/domain"
)

type CreateAudioVariantInput struct {
	TrackID      string
	Codec        domain.Codec
	BitrateKbps  int
	SampleRateHz int
	Channels     int
	DurationMs   int64
	FileURL      string
}

type CreateAudioVariantOutput struct {
	AudioVariantID string
}

type CreateAudioVariantService interface {
	Execute(ctx context.Context, input CreateAudioVariantInput) (CreateAudioVariantOutput, error)
}
