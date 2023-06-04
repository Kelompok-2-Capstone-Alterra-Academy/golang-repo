package service

import (
	"context"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/labstack/echo/v4"
)

func UploadToS3(c echo.Context, filename string, src multipart.File) (string, error) {
	logger := c.Logger()
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		logger.Fatal(err)
		return "", err
	}
	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("lms-project-capstone"),
		Key:    aws.String(filename),
		Body:   src,
	})
	if err != nil {
		logger.Fatal(err)
		return "", err
	}
	return result.Location, nil
}
