package restapi_test

import (
	"testing"

	"github.com/ariefsam/telemarketing/entity"
)

func TestCallNotHisCustomer(t *testing.T) {
	api, mockUsecase := setupAPIAndUsecase()
	dummyToken, _ := setupMockParseToken(mockUsecase)
	request := map[string]interface{}{
		"Token":       dummyToken,
		"PhoneNumber": "12312",
	}

	getCustomer := entity.Customer{
		PhoneNumber:    "12312",
		TelemarketerID: "badidtele",
	}
	filterCustomer := entity.FilterCustomer{
		PhoneNumber: &getCustomer.PhoneNumber,
	}
	mockUsecase.On("GetCustomer", filterCustomer, 1).Return([]entity.Customer{getCustomer}, nil)

	expectedResponse := map[string]interface{}{
		"Error": "Forbidden Customer",
	}

	assertRequestResponse(t, api.Call, request, expectedResponse)

}

func TestCallHisCustomer(t *testing.T) {
	api, mockUsecase := setupAPIAndUsecase()
	dummyToken, telemarketer := setupMockParseToken(mockUsecase)
	status := "Tertarik"
	request := map[string]interface{}{
		"Token":       dummyToken,
		"PhoneNumber": "12312",
		"Status":      "Tertarik",
	}

	getCustomer := entity.Customer{
		PhoneNumber:    "12312",
		TelemarketerID: telemarketer.ID,
	}
	filterCustomer := entity.FilterCustomer{
		PhoneNumber: &getCustomer.PhoneNumber,
	}
	mockUsecase.On("GetCustomer", filterCustomer, 1).Return([]entity.Customer{getCustomer}, nil)

	currentTimestamp := int64(1500000000000)
	mockUsecase.On("CurrentTimestamp").Return(currentTimestamp)

	callLog := entity.CallLog{
		Name:           getCustomer.Name,
		PhoneNumber:    getCustomer.PhoneNumber,
		TelemarketerID: telemarketer.ID,
		Status:         status,
		Timestamp:      currentTimestamp,
	}

	savedCustomer := getCustomer
	savedCustomer.Status = status
	savedCustomer.TimestampUpdated = currentTimestamp

	mockUsecase.On("SaveCallLog", callLog).Return(nil)
	mockUsecase.On("SaveCustomer", savedCustomer).Return(nil)

	expectedResponse := map[string]interface{}{
		"CallLog":  callLog,
		"Customer": savedCustomer,
	}

	assertRequestResponse(t, api.Call, request, expectedResponse)
	mockUsecase.AssertCalled(t, "SaveCallLog", callLog)
	mockUsecase.AssertCalled(t, "SaveCustomer", savedCustomer)

}
