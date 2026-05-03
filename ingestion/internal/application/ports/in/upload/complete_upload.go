package in

import (
	"context"
)

type CompleteUploadInput struct {
	UploadID  string
}

type CompleteUploadOutput struct {}

type CompleteUploadService interface{
	Execute(ctx context.Context, input CompleteUploadInput) (CompleteUploadOutput, error)
}