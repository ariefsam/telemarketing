package firebaseauth

import (
	"context"

	"github.com/ariefsam/telemarketing/entity"
	"github.com/pkg/errors"
)

type FirebaseAuth struct {
	CredentialPath string
}

func (m *FirebaseAuth) ValidateIDToken(idToken string) (authData entity.AuthData, isValid bool, err error) {

	ctx := context.TODO()
	app, err := m.getApp()
	if err != nil {
		err = errors.Wrap(err, "Firebase initiated failed")
		return
	}
	client, err := app.Auth(ctx)

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		err = errors.Wrap(err, "Error verify id token firebase")
		return
	}

	authData.Email = token.Claims["email"].(string)
	if _, ok := token.Claims["name"]; ok {
		authData.Name = token.Claims["name"].(string)
	}

	if _, ok := token.Claims["picture"]; ok {
		authData.Picture = token.Claims["picture"].(string)
	}
	isValid = true
	return
}
