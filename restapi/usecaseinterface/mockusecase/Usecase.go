package mockusecase

import (
	"github.com/ariefsam/telemarketing/entity"
	"github.com/stretchr/testify/mock"
)

type Usecase struct {
	mock.Mock
}

func (m *Usecase) AssignCustomer(telemarketerID string) (customer entity.Customer, err error) {
	args := m.Called(telemarketerID)
	customer = args.Get(0).(entity.Customer)
	err = args.Error(1)
	return
}

func (m *Usecase) GetCallLog(filter entity.FilterCallLog, limit int) (callLogs []entity.CallLog, err error) {
	args := m.Called(filter, limit)
	callLogs = args.Get(0).([]entity.CallLog)
	err = args.Error(1)
	return
}

func (m *Usecase) GetCustomer(filter entity.FilterCustomer, limit int) (customers []entity.Customer, err error) {
	args := m.Called(filter, limit)
	customers = args.Get(0).([]entity.Customer)
	err = args.Error(1)
	return
}

func (m *Usecase) SaveCallLog(callLog entity.CallLog) (err error) {
	args := m.Called(callLog)
	err = args.Error(0)
	return
}

func (m *Usecase) SaveCustomer(customer entity.Customer) (err error) {
	args := m.Called(customer)
	err = args.Error(0)
	return
}

func (m *Usecase) ParseToken(token string) (isValid bool, telemarketer entity.Telemarketer) {
	args := m.Called(token)
	isValid = args.Bool(0)
	telemarketer = args.Get(1).(entity.Telemarketer)
	return
}

func (m *Usecase) LoginByFirebase(firebaseToken string) (token string, telemarketer entity.Telemarketer, isValid bool, err error) {
	args := m.Called(firebaseToken)
	token = args.String(0)
	telemarketer = args.Get(1).(entity.Telemarketer)
	isValid = args.Bool(2)
	err = args.Error(3)
	return
}

func (m *Usecase) CurrentTimestamp() (timestamp int64) {
	args := m.Called()
	timestamp = args.Get(0).(int64)
	return
}

func (m *Usecase) SaveTelemarketer(telemarketer entity.Telemarketer) (err error) {
	args := m.Called(telemarketer)
	err = args.Error(0)
	return
}

func (m *Usecase) CreateTelemarketer(telemarketer entity.Telemarketer) (err error) {
	args := m.Called(telemarketer)
	err = args.Error(0)
	return
}

func (m *Usecase) GetTelemarketer(filter entity.FilterTelemarketer, limit int) (telemarketers []entity.Telemarketer, err error) {
	args := m.Called(filter, limit)
	telemarketers = args.Get(0).([]entity.Telemarketer)
	err = args.Error(1)
	return
}
