package firestore

import (
	"cloud.google.com/go/storage"
)

type FireStore struct {
	name   string
	Bucket *storage.BucketHandle
}
