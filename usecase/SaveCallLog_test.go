package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariefsam/telemarketing/entity"
	"github.com/ariefsam/telemarketing/usecase"
	"github.com/ariefsam/telemarketing/usecase/dependency/mockdependency"
)

func TestSaveCallLog(t *testing.T) {
	var u usecase.Usecase
	callLog := entity.CallLog{
		PhoneNumber:       "0123123",
		Status:            "Terdaftar",
		TelemarketerEmail: "t01@gmail.com",
	}
	var mockCallLogRepository mockdependency.CallLogRepository
	mockCallLogRepository.On("Save", callLog).Return(nil)
	u.CallLogRepository = &mockCallLogRepository
	err := u.SaveCallLog(callLog)
	assert.NoError(t, err)
	mockCallLogRepository.AssertCalled(t, "Save", callLog)

}
