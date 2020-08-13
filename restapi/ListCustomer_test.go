package restapi_test

import (
	"testing"

	"github.com/ariefsam/telemarketing/entity"
	"github.com/ariefsam/telemarketing/restapi"
	"github.com/ariefsam/telemarketing/restapi/usecaseinterface/mockusecase"
)

func TestListCustomerNotAdmin(t *testing.T) {
	var api restapi.RestAPI

	var mockUsecase mockusecase.Usecase
	api.Usecase = &mockUsecase

	dummyToken := "dummyTokenxxx"
	limit := 100
	request := map[string]interface{}{
		"Token": dummyToken,
		"Limit": limit,
	}

	expectedTelemarketer := entity.Telemarketer{
		Email:   "arief@gmail.com",
		IsAdmin: false,
	}

	mockUsecase.On("ParseToken", dummyToken).Return(true, expectedTelemarketer)

	filter := entity.FilterCustomer{
		TelemarketerID: &expectedTelemarketer.Email,
	}

	expectedCustomers := []entity.Customer{
		entity.Customer{
			Name:        "A",
			PhoneNumber: "09123",
		},
		entity.Customer{
			Name:        "B",
			PhoneNumber: "091232",
		},
	}
	mockUsecase.On("GetCustomer", filter, limit).Return(expectedCustomers, nil)

	expectedResponse := map[string]interface{}{
		"Customers": expectedCustomers,
	}

	assertRequestResponse(t, api.ListCustomer, request, expectedResponse)
}

func TestListCustomerNotAdminBadTelemarketerID(t *testing.T) {
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
		TelemarketerID: &expectedTelemarketer.Email,
	}
	expectedCustomers := dummyCustomers()
	mockUsecase.On("GetCustomer", filter, limit).Return(expectedCustomers, nil)

	expectedResponse := map[string]interface{}{
		"Customers": expectedCustomers,
	}

	assertRequestResponse(t, api.ListCustomer, request, expectedResponse)
}

func dummyTelemarketerNotAdmin() (expectedTelemarketer entity.Telemarketer) {
	expectedTelemarketer = entity.Telemarketer{
		Email:   "arief@gmail.com",
		IsAdmin: false,
	}
	return
}

func dummyCustomers() []entity.Customer {
	return []entity.Customer{
		entity.Customer{
			Name:        "A",
			PhoneNumber: "09123",
		},
		entity.Customer{
			Name:        "B",
			PhoneNumber: "091232",
		},
	}
}
