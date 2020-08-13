package firebaseauth

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func (f *FirebaseAuth) getApp() (app *firebase.App, err error) {
	opt := option.WithCredentialsFile(f.CredentialPath)
	app, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
