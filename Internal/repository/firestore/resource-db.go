package firestore

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"time"

	"cloud.google.com/go/storage"
	"golang.org/x/oauth2/jwt"
)

func (f FireStore) UploadImage(ctx context.Context, filename string, file multipart.File) (string, error) {

	var (
		url string
	)

	wc := f.Bucket.Object(filename).NewWriter(ctx)
	_, err := io.Copy(wc, file)
	if err != nil {
		log.Println("[upload domain][UploadImage] write err,", err)
		return url, err
	}

	if err := wc.Close(); err != nil {
		log.Println("[upload domain][UploadImage] writer close err,", err)
		return url, err
	}

	url = "https://storage.cloud.google.com/" + f.name + "/" + filename

	return url, nil
}

// generateV4PutObjectSignedURL generates object signed URL with PUT method.
func (f FireStore) GenerateV4PutObjectSignedURL(ctx context.Context, bucket, object string, config *jwt.Config) (string, error) {

	opts := &storage.SignedURLOptions{
		Scheme:         storage.SigningSchemeV4,
		Method:         "GET",
		GoogleAccessID: config.Email,
		PrivateKey:     config.PrivateKey,
		Expires:        time.Now().Add(15 * time.Minute),
	}

	u, err := storage.SignedURL(bucket, object, opts)
	if err != nil {
		return "", fmt.Errorf("storage.SignedURL: %v", err)
	}

	return u, nil
}
