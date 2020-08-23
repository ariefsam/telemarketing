package restapi_test

import (
	"testing"

	"github.com/ariefsam/telemarketing/entity"
)

func TestGetCallLogNotAdmin(t *testing.T) {
	api, mockUsecase := setupAPIAndUsecase()
	dummyToken, telemarketer := setupMockParseToken(mockUsecase)
	request := map[string]interface{}{
		"Token": dummyToken,
		"FilterCallLog": map[string]interface{}{
			"Status": "Tertarik",
		},
		"Limit": 1,
	}
	status := "Tertarik"
	filter := entity.FilterCallLog{
		TelemarketerEmail: &telemarketer.Email,
		Status:            &status,
	}
	limit := 1
	callLogs := []entity.CallLog{
		entity.CallLog{
			Name:        "c",
			PhoneNumber: "082323",
		},
		entity.CallLog{
			Name:        "a",
			PhoneNumber: "012382323",
		},
	}
	mockUsecase.On("GetCallLog", filter, limit).Return(callLogs, nil)

	response := map[string]interface{}{
		"CallLogs": callLogs,
	}
	assertRequestResponse(t, api.GetCallLog, request, response)
}
