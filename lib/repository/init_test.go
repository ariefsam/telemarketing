package repository_test

import (
	"github.com/ariefsam/telemarketing/configuration"
	"github.com/ariefsam/telemarketing/lib/repository"
)

func init() {
	repository.FirebaseAccountPath = configuration.FirebaseCredentialPath
}
