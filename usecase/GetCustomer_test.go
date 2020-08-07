package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariefsam/telemarketing/entity"
	"github.com/ariefsam/telemarketing/usecase"
	"github.com/ariefsam/telemarketing/usecase/dependency/mockdependency"
)

func TestGetCustomer(t *testing.T) {
	var u usecase.Usecase
	filter := entity.FilterCustomer{
		TelemarketerID: "aaa",
	}
	limit := 10
	expectedCustomer := []entity.Customer{
		entity.Customer{
			PhoneNumber:    "0343",
			TelemarketerID: "aaa",
			Status:         "Empty",
		},
		entity.Customer{
			PhoneNumber:    "03431",
			TelemarketerID: "aaa",
			Status:         "Empty",
		},
	}
	var mockCustomerRepository mockdependency.CustomerRepository
	u.CustomerRepository = &mockCustomerRepository
	mockCustomerRepository.On("Get", filter, limit).Return(expectedCustomer, nil)
	customers, err := u.GetCustomer(filter, limit)
	assert.NoError(t, err)
	assert.Equal(t, expectedCustomer, customers)
}
