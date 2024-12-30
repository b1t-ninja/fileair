package aws3

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
	"time"
)

func UploadToS3AndGeneratePresignedURL(path string, expiration time.Duration) (string, error) {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		return "", fmt.Errorf("error loading .env file: %v", err)
	}

	bucketName := os.Getenv("AWS_BUCKET_NAME")
	objectKey := filepath.Base(path) // Use the file name as the object key

	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", fmt.Errorf("failed to load AWS config: %v", err)
	}

	// Open the file
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Upload file to S3
	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)
	_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %v", err)
	}

	// Generate a presigned URL
	presignClient := s3.NewPresignClient(client)
	presignRequest, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}, s3.WithPresignExpires(expiration))
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL: %v", err)
		os.Exit(1)
	}

	return presignRequest.URL, nil
}
