package domain

import (
	"context"
	"mime/multipart"
)

type S3Service interface {
	UploadFiles(ctx context.Context, files []*multipart.FileHeader, folder string) ([]string, error)
}
