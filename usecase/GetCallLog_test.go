package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariefsam/telemarketing/entity"
	"github.com/ariefsam/telemarketing/usecase"
	"github.com/ariefsam/telemarketing/usecase/dependency/mockdependency"
)

func TestGetCallLog(t *testing.T) {
	var u usecase.Usecase
	filter := entity.FilterCallLog{
		TimestampEnd: 1300000000,
	}
	limit := 10
	expectedCallLogs := []entity.CallLog{
		entity.CallLog{
			TelemarketerEmail: "s@gmail.com",
			PhoneNumber:       "09234",
			Timestamp:         123000,
		},
		entity.CallLog{
			TelemarketerEmail: "s@gmail.com",
			PhoneNumber:       "09234",
			Timestamp:         123100,
		},
	}

	var mockCallLogRepository mockdependency.CallLogRepository
	u.CallLogRepository = &mockCallLogRepository

	mockCallLogRepository.On("Get", filter, limit).Return(expectedCallLogs, nil)

	callLogs, err := u.GetCallLog(filter, limit)
	assert.NoError(t, err)
	assert.Equal(t, expectedCallLogs, callLogs)

}
