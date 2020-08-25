package restapi_test

import (
	"testing"

	"github.com/ariefsam/telemarketing/entity"
)

func TestListCustomerNotAdmin(t *testing.T) {
	api, mockUsecase := setupAPIAndUsecase()
	dummyToken, expectedTelemarketer := setupMockParseToken(mockUsecase)
	limit := 100

	request := map[string]interface{}{
		"Token": dummyToken,
		"Limit": limit,
	}

	filter := entity.FilterCustomer{
		TelemarketerID: &expectedTelemarketer.ID,
	}

	expectedCustomers := dummyCustomers()
	mockUsecase.On("GetCustomer", filter, limit).Return(expectedCustomers, nil)

	expectedResponse := map[string]interface{}{
		"Customers": expectedCustomers,
	}

	assertRequestResponse(t, api.ListCustomer, request, expectedResponse)
}

func TestListCustomerNotAdminBadTelemarketerID(t *testing.T) {
	api, mockUsecase := setupAPIAndUsecase()
	dummyToken, _ := setupMockParseToken(mockUsecase)
	limit := 100

	request := map[string]interface{}{
		"Token": dummyToken,
		"Limit": limit,
		"FilterCustomer": map[string]interface{}{
			"TelemarketerID": "xxx",
		},
	}

	expectedResponse := map[string]interface{}{
		"Error": "FilterCustomer TelemarketerID forbidden",
	}

	assertRequestResponse(t, api.ListCustomer, request, expectedResponse)
}
