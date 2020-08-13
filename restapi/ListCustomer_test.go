package restapi_test

import (
	"testing"

	"github.com/ariefsam/telemarketing/entity"
	"github.com/ariefsam/telemarketing/restapi"
	"github.com/ariefsam/telemarketing/restapi/usecaseinterface/mockusecase"
)

func TestListCustomerAdmin(t *testing.T) {
	var api restapi.RestAPI
	dummyToken := "dummyTokenxxx"
	request := map[string]interface{}{
		"Token": dummyToken,
	}

	expectedTelemarketer := entity.Telemarketer{
		Email: "arief@gmail.com",
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

	expectedResponse := map[string]interface{}{
		"Customers": expectedCustomers,
	}

	var mockUsecase mockusecase.Usecase
	api.Usecase = &mockUsecase

	mockUsecase.On("ParseToken", dummyToken).Return(true, expectedTelemarketer)

	assertRequestResponse(t, api.ListCustomer, request, expectedResponse)
}
