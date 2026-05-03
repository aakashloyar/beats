package out

import (
	"context"
	"github.com/aakashloyar/beats/ingestion/internal/domain"
)

type Storage interface {
	CreateMultipartUpload(ctx context.Context, key string) (string, error)
	GeneratePresignedPartURL(ctx context.Context, key string, uploadID string, partNumber int32) (string, error)
	CompleteMultipartUpload(ctx context.Context, key string, uploadID string, parts []domain.Chunk) error
	BuildStorageKey(uploadID, fileName string) string
}
