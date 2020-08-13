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
	teleID := "aaa"
	filter := entity.FilterCustomer{
		TelemarketerEmail: &teleID,
	}
	limit := 10
	expectedCustomer := []entity.Customer{
		entity.Customer{
			PhoneNumber:       "0343",
			TelemarketerEmail: "aaa",
			Status:            "Empty",
		},
		entity.Customer{
			PhoneNumber:       "03431",
			TelemarketerEmail: "aaa",
			Status:            "Empty",
		},
	}
	var mockCustomerRepository mockdependency.CustomerRepository
	u.CustomerRepository = &mockCustomerRepository
	mockCustomerRepository.On("Get", filter, limit).Return(expectedCustomer, nil)
	customers, err := u.GetCustomer(filter, limit)
	assert.NoError(t, err)
	assert.Equal(t, expectedCustomer, customers)
}
