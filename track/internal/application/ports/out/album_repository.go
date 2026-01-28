package out

import (
	"github.com/aakashloyar/beats/track/internal/domain"
)

type AlbumRepository interface {
	Save(domain.Album) error
	FindByID(albumID string) (domain.Album, error)
	ListAlbums(title string) ([]domain.Album, error)
}
