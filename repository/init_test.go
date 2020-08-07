package repository_test

import (
	"github.com/ariefsam/telemarketing/configuration"
	"github.com/ariefsam/telemarketing/repository"
)

func init() {
	repository.FirebaseAccountPath = configuration.FirebaseCredentialPath
}
