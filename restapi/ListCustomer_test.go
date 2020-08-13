package restapi_test

import (
	"testing"

	"github.com/ariefsam/telemarketing/entity"
	"github.com/ariefsam/telemarketing/restapi"
	"github.com/ariefsam/telemarketing/restapi/usecaseinterface/mockusecase"
)

func TestListCustomerNotAdmin(t *testing.T) {
	dummyToken := "dummyTokenxxx"
	limit := 100

	request := map[string]interface{}{
		"Token": dummyToken,
		"Limit": limit,
	}

	var api restapi.RestAPI
	var mockUsecase mockusecase.Usecase
	api.Usecase = &mockUsecase

	expectedTelemarketer := dummyTelemarketerNotAdmin()
	mockUsecase.On("ParseToken", dummyToken).Return(true, expectedTelemarketer)
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
	dummyToken := "dummyTokenxxx"
	limit := 100

	request := map[string]interface{}{
		"Token": dummyToken,
		"Limit": limit,
		"FilterCustomer": map[string]interface{}{
			"TelemarketerEmail": "xxx",
		},
	}

	var api restapi.RestAPI
	var mockUsecase mockusecase.Usecase
	api.Usecase = &mockUsecase

	expectedTelemarketer := dummyTelemarketerNotAdmin()
	mockUsecase.On("ParseToken", dummyToken).Return(true, expectedTelemarketer)

	expectedResponse := map[string]interface{}{
		"Error": "FilterCustomer TelemarketerEmail forbidden",
	}

	assertRequestResponse(t, api.ListCustomer, request, expectedResponse)
}
