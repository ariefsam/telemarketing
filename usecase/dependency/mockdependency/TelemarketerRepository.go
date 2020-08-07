package mockdependency

import (
	"github.com/ariefsam/telemarketing/entity"
	"github.com/stretchr/testify/mock"
)

type TelemarketerRepository struct {
	mock.Mock
}

func (m *TelemarketerRepository) Save(telemarketer entity.Telemarketer) (err error) {
	args := m.Called(telemarketer)
	err = args.Error(0)
	return
}

func (m *TelemarketerRepository) Get(filter entity.FilterTelemarketer, limit int) (telemarketers []entity.Telemarketer, err error) {
	args := m.Called(filter, limit)
	telemarketers = args.Get(0).([]entity.Telemarketer)
	err = args.Error(1)
	return
}
