package postgres

import (
	"database/sql"

	"github.com/aakashloyar/beats/track/internal/application/ports/out"
	"github.com/aakashloyar/beats/track/internal/domain"
)

type AudioVariantRepository struct {
	db *sql.DB
}

func NewAudioVariantRepository(db *sql.DB) out.AudioVariantRepository {
	return &AudioVariantRepository{db: db}
}

func (r *AudioVariantRepository) Save(v domain.AudioVariant) error {
	query := `
		INSERT INTO audio_variants (
			id,
			track_id,
			codec,
			bitrate_kbps,
			sample_rate_hz,
			channels,
			duration_ms,
			file_url,
			created_at
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
	`

	_, err := r.db.Exec(
		query,
		v.ID,
		v.TrackID,
		v.Codec,
		v.BitrateKbps,
		v.SampleRateHz,
		v.Channels,
		v.DurationMs,
		v.FileURL,
		v.CreatedAt,
	)

	return err
}
