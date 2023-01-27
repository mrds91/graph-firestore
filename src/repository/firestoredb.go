package repository

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
)

var fireStoreClient *firestore.Client

func CreateFirestoreClient() *firestore.Client {
	// Sets your Google Cloud Platform project ID.
	if fireStoreClient != nil {
		return fireStoreClient
	}
	projectID := "first-project-375903"

	// [END firestore_setup_client_create]
	// Override with -project flags
	//	flag.StringVar(&projectID, "project", projectID, "The Google Cloud Platform project ID.")
	//	flag.Parse()

	ctx := context.Background()
	// [START firestore_setup_client_create]
	fmt.Println("Creating client...")
	client, err := firestore.NewClient(ctx, projectID)
	fmt.Println("New client...", err)
	//	defer client.Close()
	if err != nil {
		fmt.Println("Failed to create client: %v", err)
	}
	// Close client when done with
	// defer client.Close()
	fireStoreClient = client
	return client
}
