package main

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/joho/godotenv"
	"github.com/keimoon/gore"
	"google.golang.org/api/option"
)

var ProjectionDB string
var FirebaseAccountPath string

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ProjectionDB = os.Getenv("projection_db")
	FirebaseAccountPath = "../firebase.json"
	conn, err := gore.Dial("localhost:6379") //Connect to redis server at localhost:6379
	gore.NewCommand("SELECT", 15).Run(conn)
	gore.NewCommand("FLUSHDB").Run(conn)
	if err != nil {
		return
	}
	defer conn.Close()
	ctx, docRef, err := getFirestoreClient()
	if err != nil {
		log.Println(err)
		return
	}
	docRef.Delete(ctx)

	_, err = gore.NewCommand("DEL", "event-from-server-last-time").Run(conn)
	if err != nil {
		log.Println(err)
	}
}

func getFirestoreClient() (ctx context.Context, docRef *firestore.DocumentRef, err error) {
	if ProjectionDB == "" {
		ProjectionDB = "telemarketing"
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
	docRef = client.Collection("db").Doc(ProjectionDB)
	return
}
