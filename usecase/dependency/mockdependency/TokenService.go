package mockdependency

import (
	"github.com/ariefsam/telemarketing/entity"
	"github.com/stretchr/testify/mock"
)

type TokenService struct {
	mock.Mock
}

func (m *TokenService) Create(telemarketer entity.Telemarketer) (token string) {
	args := m.Called(telemarketer)
	token = args.String(0)
	return
}
func (m *TokenService) Parse(token string) (isValid bool, telemarketer entity.Telemarketer) {
	args := m.Called(token)
	isValid = args.Bool(0)
	telemarketer = args.Get(1).(entity.Telemarketer)
	return
}
