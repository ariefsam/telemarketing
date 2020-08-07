package ioc_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariefsam/telemarketing/firebaseauth"
	"github.com/ariefsam/telemarketing/ioc"
	"github.com/ariefsam/telemarketing/repository"
)

func TestUsecase(t *testing.T) {
	u := ioc.Usecase()
	assert.NotNil(t, u.CustomerRepository)
	assert.NotNil(t, u.CallLogRepository)
	assert.NotEmpty(t, repository.FirebaseAccountPath)
	assert.NotNil(t, u.TokenService)
	assert.NotNil(t, u.FirebaseAuth)
	if u.FirebaseAuth != nil {
		f := u.FirebaseAuth.(*firebaseauth.FirebaseAuth)
		assert.NotEmpty(t, f.CredentialPath)
	}
	assert.NotNil(t, u.TelemarketerRepository)

}
