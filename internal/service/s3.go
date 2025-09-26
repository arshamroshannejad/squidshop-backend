package service

import (
	"context"
	"fmt"
	"log/slog"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/arshamroshannejad/squidshop-backend/config"
	"github.com/arshamroshannejad/squidshop-backend/internal/domain"
	"github.com/aws/aws-sdk-go-v2/aws"
	cfgAws "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type s3ServiceImpl struct {
	client *s3.Client
	bucket string
	region string
	logger *slog.Logger
}

func NewS3Service(cfg *config.Config, logger *slog.Logger) domain.S3Service {
	s3Config, err := cfgAws.LoadDefaultConfig(context.Background(), cfgAws.WithRegion(cfg.S3.Region))
	if err != nil {
		logger.Error("failed to load default config", "error", err)
		return nil
	}
	s3Config.Credentials = aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
		return aws.Credentials{
			AccessKeyID:     cfg.S3.AccessKey,
			SecretAccessKey: cfg.S3.SecretKey,
		}, nil
	})
	s3Config.BaseEndpoint = aws.String(cfg.S3.Endpoint)
	client := s3.NewFromConfig(s3Config)
	return &s3ServiceImpl{
		client: client,
		bucket: cfg.S3.Bucket,
		region: cfg.S3.Region,
		logger: logger,
	}
}

func (s *s3ServiceImpl) UploadFiles(ctx context.Context, files []*multipart.FileHeader, folder string) ([]string, error) {
	var uploadedImages []string
	for _, file := range files {
		fileBytes, err := file.Open()
		if err != nil {
			s.logger.Error("failed to open file", "error", err)
			return uploadedImages, err
		}
		ext := filepath.Ext(file.Filename)
		name := strings.TrimSuffix(file.Filename, ext)
		name = strings.ReplaceAll(name, " ", "_")
		ts := time.Now().UnixNano()
		key := fmt.Sprintf("%s/%s_%d%s", folder, name, ts, ext)
		_, err = s.client.PutObject(ctx, &s3.PutObjectInput{
			Bucket: aws.String(s.bucket),
			Key:    aws.String(key),
			Body:   fileBytes,
		})
		fileBytes.Close()
		if err != nil {
			s.logger.Error("failed to upload file", "error", err)
			return uploadedImages, err
		}
		uploadedImages = append(uploadedImages, key)
	}
	return uploadedImages, nil
}
