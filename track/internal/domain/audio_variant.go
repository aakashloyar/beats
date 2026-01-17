// internal/domain/audio_variant.go
package domain

import "time"

type AudioVariant struct {
	ID           string
	TrackID      string
	Codec        Codec
	BitrateKbps  int
	SampleRateHz int
	Channels     int
	DurationMs   int64
	FileURL      string
	CreatedAt    time.Time
}
