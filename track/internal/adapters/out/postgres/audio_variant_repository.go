package postgres

import (
	"database/sql"

	"github.com/aakashloyar/beats/track/internal/application/ports/out"
	"github.com/aakashloyar/beats/track/internal/domain"
)

type PostgresAudioVariantRepository struct {
	db *sql.DB
}

func NewAudioVariantRepository(db *sql.DB) out.AudioVariantRepository {
	return &PostgresAudioVariantRepository{db: db}
}

func (r *PostgresAudioVariantRepository) Save(v *domain.AudioVariant) error {
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

func (r *PostgresAudioVariantRepository) FindByTrackID(trackID string) ([]*domain.AudioVariant, error) {
	query := `
		SELECT
			id,
			track_id,
			codec,
			bitrate_kbps,
			sample_rate_hz,
			channels,
			duration_ms,
			file_url,
			created_at
		FROM audio_variants
		WHERE track_id = $1
		ORDER BY bitrate_kbps ASC
	`

	rows, err := r.db.Query(query, trackID)

	if err != nil {
		return nil, err 
	}
	defer rows.Close()

	var variants []*domain.AudioVariant 

	for rows.Next() {
		var v domain.AudioVariant
		err := rows.Scan(
			&v.ID,
			&v.TrackID,
			&v.Codec,
			&v.BitrateKbps,
			&v.SampleRateHz,
			&v.Channels,
			&v.DurationMs,
			&v.FileURL,
			&v.CreatedAt,
		)
		if err != nil {
			return nil, err 
		}
		variants = append(variants, &v)
	}

	if err := rows.Err(); err != nil {
		return nil, err 
	}
	return variants, nil 
}
