package services

import (
	"context"
	"fmt"
	"social_web_server/database"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type TFileUrl struct {
	PresignedUrl string `json:"presign_url"`
	FileUrl      string `json:"file_url"`
}

func (s *ServiceStruct) GetPresignedUrl(ctx context.Context, fileName string, fileType string, duration time.Duration) (*TFileUrl, error) {

	req, err := s.Presigner.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket:      &s.Bucket,
		Key:         &fileName,
		ContentType: &fileType,
	}, s3.WithPresignExpires(duration))

	if err != nil {
		return nil, fmt.Errorf("failed to generate pre-signed URL: %w", err)
	}

	return &TFileUrl{PresignedUrl: req.URL, FileUrl: fmt.Sprintf("https://%s.s3.amazonaws.com/%s", s.Bucket, fileName)}, nil
}

func (s *ServiceStruct) SaveFileUrl(ctx context.Context, params *database.CreateUploadedFileParams) error {
	if err := s.repository.InsertUploadedFile(ctx, *params); err != nil {
		return err
	}

	return nil
}
