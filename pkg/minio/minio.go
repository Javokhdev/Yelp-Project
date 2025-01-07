package minio

import (
	"context"
	"fmt"
	"log/slog"
	"mime"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/Javokhdev/Yelp-Project/config"
)

type MinioClient struct {
	Client     *minio.Client
	BucketName string
}

func NewMinioClient(cfg *config.Config) (*MinioClient, error) {
	// Initialize MinIO client using config values
	minioClient, err := minio.New(cfg.MinIO.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.MinIO.AccessKey, cfg.MinIO.SecretKey, ""),
		Secure: false,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create minio client: %v", err)
	}

	client := &MinioClient{
		Client:     minioClient,
		BucketName: cfg.MinIO.BucketName,
	}

	// Create bucket if it doesn't exist
	err = client.createBucket()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (mc *MinioClient) createBucket() error {
	exists, err := mc.Client.BucketExists(context.Background(), mc.BucketName)
	if err != nil {
		return fmt.Errorf("failed to check if bucket exists: %v", err)
	}
	if !exists {
		err = mc.Client.MakeBucket(context.Background(), mc.BucketName, minio.MakeBucketOptions{})
		if err != nil {
			return fmt.Errorf("failed to create bucket: %v", err)
		}
		slog.Info("Bucket %s created successfully", mc.BucketName, "")
	}

	policy := fmt.Sprintf(`{
		"Version": "2012-10-17",
		"Statement": [
			{
				"Effect": "Allow",
				"Principal": "*",
				"Action": ["s3:GetObject"],
				"Resource": ["arn:aws:s3:::%s/*"]
			}
		]
	}`, mc.BucketName)

	err = mc.Client.SetBucketPolicy(context.Background(), mc.BucketName, policy)
	if err != nil {
		slog.Error("Error while setting bucket policy: %v", "err", err)
		return nil
	}

	return nil
}

func (mc *MinioClient) UploadFile(ctx context.Context, fileName string, filePath string) (string, error) {
	ext := filepath.Ext(fileName)
	uuid := uuid.New().String()
	fileName = uuid + "_" + fileName
	contentType := mime.TypeByExtension(ext)

	if contentType == "" {
		contentType = "application/octet-stream"
	}

	_, err := mc.Client.FPutObject(ctx, mc.BucketName, fileName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %v", err)
	}

	serverHost := "medias"
	domain := "booknest.uz"
	minioURL := fmt.Sprintf("https://%s.%s/%s/%s", serverHost, domain, mc.BucketName, fileName)

	return minioURL, nil
}
