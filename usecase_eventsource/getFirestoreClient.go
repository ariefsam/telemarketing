package usecase_eventsource

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var FirebaseAccountPath string
var EventSourceDB string

func getFirestoreClient() (ctx context.Context, docRef *firestore.DocumentRef, err error) {
	if EventSourceDB == "" {
		EventSourceDB = "telemarketing"
	}
	if FirebaseAccountPath == "" {
		FirebaseAccountPath = "firebase.json"
	}
	ctx = context.Background()
	sa := option.WithCredentialsFile(FirebaseAccountPath)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := app.Firestore(ctx)
	docRef = client.Collection("event-source").Doc(EventSourceDB)
	return
}
