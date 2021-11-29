package util

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func InsertIntoCloudStorage() {

	config := &firebase.Config{
		StorageBucket: "gs://storage-1bb41.appspot.com",
	}

	opt := option.WithCredentialsFile("files/storage-1bb41-firebase-adminsdk-esal9-7d4133437b.json")
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Created bucket handle: %v\n", bucket)

	//TODO : read and Write image
}
