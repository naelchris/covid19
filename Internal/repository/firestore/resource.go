package firestore

import (
	"context"
	"log"
	"mime/multipart"

	"cloud.google.com/go/storage"
	"golang.org/x/oauth2/jwt"
)

type DomainItf interface {
	Upload(ctx context.Context, fileName string, file multipart.File) (string, error)
	GenerateV4PutObjectSignedURL(ctx context.Context, bucket, object string, config *jwt.Config) (string, error)
}

type DomianResouceItf interface {
	UploadImage(ctx context.Context, filename string, file multipart.File) (string, error)
	GenerateV4PutObjectSignedURL(ctx context.Context, bucket, object string, config *jwt.Config) (string, error)
}

type Domain struct {
	firestore DomianResouceItf
}

func InitDomain(bucketName string, bucket *storage.BucketHandle) Domain {
	return Domain{
		firestore: FireStore{
			name:   bucketName,
			Bucket: bucket,
		},
	}
}

func (d Domain) Upload(ctx context.Context, fileName string, file multipart.File) (string, error) {
	var (
		imageUrl string
		err      error
	)

	imageUrl, err = d.firestore.UploadImage(ctx, fileName, file)
	if err != nil {
		log.Println("[resource Upload Image][UploadImage][Upload] err", err)
		return imageUrl, err
	}

	return imageUrl, nil
}

func (d Domain) GenerateV4PutObjectSignedURL(ctx context.Context, bucket, object string, config *jwt.Config) (resp string, err error) {
	resp, err = d.firestore.GenerateV4PutObjectSignedURL(ctx, bucket, object, config)
	if err != nil {
		log.Println("[GenerateV4PutObjectSignedURL] err", err)
		return resp, err
	}

	return resp, err
}
