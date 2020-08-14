package restapi_test

import (
	"testing"

	"github.com/ariefsam/telemarketing/entity"
	"github.com/ariefsam/telemarketing/restapi"
	"github.com/ariefsam/telemarketing/restapi/usecaseinterface/mockusecase"
)

func TestListCustomerNotAdmin(t *testing.T) {
	api, mockUsecase := initAPIAndUsecase()
	dummyToken, expectedTelemarketer := setupMockParseToken(mockUsecase)
	limit := 100

	request := map[string]interface{}{
		"Token": dummyToken,
		"Limit": limit,
	}

	filter := entity.FilterCustomer{
		TelemarketerEmail: &expectedTelemarketer.Email,
	}

	expectedCustomers := dummyCustomers()
	mockUsecase.On("GetCustomer", filter, limit).Return(expectedCustomers, nil)

	expectedResponse := map[string]interface{}{
		"Customers": expectedCustomers,
	}

	assertRequestResponse(t, api.ListCustomer, request, expectedResponse)
}

func TestListCustomerNotAdminBadTelemarketerEmail(t *testing.T) {
	api, mockUsecase := initAPIAndUsecase()
	dummyToken, _ := setupMockParseToken(mockUsecase)
	limit := 100

	request := map[string]interface{}{
		"Token": dummyToken,
		"Limit": limit,
		"FilterCustomer": map[string]interface{}{
			"TelemarketerEmail": "xxx",
		},
	}

	expectedResponse := map[string]interface{}{
		"Error": "FilterCustomer TelemarketerEmail forbidden",
	}

	assertRequestResponse(t, api.ListCustomer, request, expectedResponse)
}

func initAPIAndUsecase() (api *restapi.RestAPI, mockUsecase *mockusecase.Usecase) {
	var m mockusecase.Usecase
	api = new(restapi.RestAPI)
	mockUsecase = &m
	api.Usecase = mockUsecase
	return
}
