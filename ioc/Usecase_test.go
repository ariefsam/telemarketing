package ioc_test

import (
	"testing"

	"github.com/ariefsam/telemarketing/ioc"
	"github.com/ariefsam/telemarketing/lib/firebaseauth"
	"github.com/ariefsam/telemarketing/lib/repository"
	"github.com/stretchr/testify/assert"
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
	assert.NotNil(t, u.Timer)
	assert.NotNil(t, u.IDGenerator)
	assert.NotNil(t, u.ReportRepository)

}
