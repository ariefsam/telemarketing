package restapi_test

import (
	"testing"
)

func TestAssignCustomer(t *testing.T) {
	api, mockUsecase := setupAPIAndUsecase()

	dummyToken, expectedTelemarketer := setupMockParseToken(mockUsecase)
	request := map[string]interface{}{
		"Token": dummyToken,
	}

	expectedCustomer := dummyCustomers()[0]
	expectedCustomer.TelemarketerID = expectedTelemarketer.ID
	mockUsecase.On("AssignCustomer", expectedTelemarketer.ID).Return(expectedCustomer, nil)

	expectedResponse := map[string]interface{}{
		"Customer": expectedCustomer,
	}
	assertRequestResponse(t, api.AssignCustomer, request, expectedResponse)

}
