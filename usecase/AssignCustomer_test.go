package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariefsam/telemarketing/entity"
	"github.com/ariefsam/telemarketing/usecase"
	"github.com/ariefsam/telemarketing/usecase/dependency/mockdependency"
)

func TestAssignCustomer(t *testing.T) {
	var u usecase.Usecase
	telemarketerID := "tele0001"
	limit := 1

	var mockCustomerRepository mockdependency.CustomerRepository
	u.CustomerRepository = &mockCustomerRepository

	filter := entity.FilterCustomer{
		TelemarketerID: "Not Set",
	}

	expectedCustomer := entity.Customer{
		Name:        "Member01",
		PhoneNumber: "073434",
	}
	mockCustomerRepository.On("Get", filter, limit).Return([]entity.Customer{expectedCustomer}, nil)

	expectedSavedCustomer := expectedCustomer
	expectedSavedCustomer.TelemarketerID = telemarketerID
	mockCustomerRepository.On("Save", expectedSavedCustomer).Return(nil)

	customer, err := u.AssignCustomer(telemarketerID)
	assert.NoError(t, err)
	assert.Equal(t, expectedSavedCustomer, customer)
	mockCustomerRepository.AssertCalled(t, "Save", expectedSavedCustomer)

}
