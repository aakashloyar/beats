package out
import "github.com/aakashloyar/beats/track/internal/domain"
type TrackRespository interface {
	Save(track *domain.Track) error
	FindById(id string) (*domain.Track, error)
	FindByArtist(artistID string) (*domain.Track, error)
}

