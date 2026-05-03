package out

import "context"

type UploadCompletedEvent struct {
	UploadID   string `json:"upload_id"`
	StorageKey string `json:"storage_key"`
	ArtistID   string `json:"artist_id"`
}

type EventPublisher interface {
	PublishUploadCompleted(ctx context.Context, event UploadCompletedEvent) error
}