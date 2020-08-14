package restapi_test

import (
	"testing"
)

func TestAssignCustomer(t *testing.T) {
	api, mockUsecase := initAPIAndUsecase()

	dummyToken, expectedTelemarketer := setupMockParseToken(mockUsecase)
	request := map[string]interface{}{
		"Token": dummyToken,
	}

	expectedCustomer := dummyCustomers()[0]
	expectedCustomer.TelemarketerEmail = expectedTelemarketer.Email
	mockUsecase.On("AssignCustomer", expectedTelemarketer.Email).Return(expectedCustomer, nil)

	expectedResponse := map[string]interface{}{
		"Customer": expectedCustomer,
	}
	assertRequestResponse(t, api.AssignCustomer, request, expectedResponse)

}
