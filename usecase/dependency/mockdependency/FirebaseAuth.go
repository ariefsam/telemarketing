package mockdependency

import (
	"github.com/ariefsam/telemarketing/entity"
	"github.com/stretchr/testify/mock"
)

type FirebaseAuth struct {
	mock.Mock
}

func (m *FirebaseAuth) ValidateIDToken(token string) (authData entity.AuthData, isValid bool, err error) {
	args := m.Called(token)
	isValid = args.Bool(1)
	authData = args.Get(0).(entity.AuthData)
	err = args.Error(2)
	return
}
