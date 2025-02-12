package services

import (
	"social_web_server/repository"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type ServiceStruct struct {
	repository *repository.RepositoryStruct
	Bucket     string
	S3Client   *s3.Client
}

func (s *ServiceStruct) InitialiseService(r *repository.RepositoryStruct, Bucket string, s3Client *s3.Client) *ServiceStruct {
	s.repository = r
	s.Bucket = Bucket
	s.S3Client = s3Client

	return s
}
