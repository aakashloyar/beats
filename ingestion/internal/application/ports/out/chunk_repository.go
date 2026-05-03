package out
import (
	"github.com/aakashloyar/beats/ingestion/internal/domain"
)
type ChunkRespository interface {
	Mark(chunk domain.Chunk) error
}
