package mockdependency

import (
	"github.com/ariefsam/telemarketing/entity"
	"github.com/stretchr/testify/mock"
)

type CallLogRepository struct {
	mock.Mock
}

func (m *CallLogRepository) Save(callLog entity.CallLog) (err error) {
	args := m.Called(callLog)
	err = args.Error(0)
	return
}
func (m *CallLogRepository) Get(filter entity.FilterCallLog, limit int) (callLogs []entity.CallLog, err error) {
	args := m.Called(filter, limit)
	callLogs = args.Get(0).([]entity.CallLog)
	err = args.Error(1)
	return
}
