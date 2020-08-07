package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariefsam/telemarketing/entity"
	"github.com/ariefsam/telemarketing/usecase"
	"github.com/ariefsam/telemarketing/usecase/dependency/mockdependency"
)

func TestSaveCustomer(t *testing.T) {
	var u usecase.Usecase
	customer := entity.Customer{
		PhoneNumber:    "0123123",
		Status:         "Terdaftar",
		TelemarketerID: "t01",
	}
	var mockCustomerRepository mockdependency.CustomerRepository
	u.CustomerRepository = &mockCustomerRepository
	mockCustomerRepository.On("Save", customer).Return(nil)
	err := u.SaveCustomer(customer)
	assert.NoError(t, err)
	mockCustomerRepository.AssertCalled(t, "Save", customer)
}
