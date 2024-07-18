package firestorage

import (
	"context"
	"fmt"
	"io"
	"os"

	"path/filepath"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type FireStorage interface {
	Upload(bucketName, objectName, filePath string) (string, error)
}

type fireStorageRepo struct {
}

func NewFireStorage() FireStorage {
	repo := &fireStorageRepo{}

	return repo
}

func (p *fireStorageRepo) Upload(bucketName, objectName, filePath string) (string, error) {
	ctx := context.Background()

	credentialsPath := filepath.Join(os.Getenv("FIRESTORESDK_PATH"))
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialsPath))
	if err != nil {
		return "", err
	}
	defer client.Close()

	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	wc := client.Bucket(bucketName).Object(objectName).NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		return "", err
	}
	if err := wc.Close(); err != nil {
		return "", err
	}

	attrs := wc.Attrs()

	url := attrs.MediaLink
	fmt.Printf("File %s uploaded to %s.\n", filePath, url)
	return url, nil
}
