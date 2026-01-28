package out

import (
	"github.com/aakashloyar/beats/track/internal/domain"
)

type ArtistRepository interface {
	Save(domain.Artist) error
	FindByID(artistID string) (domain.Artist, error)
}
