package usecase_test

import (
	"testing"

	"github.com/ariefsam/telemarketing/entity"
	"github.com/ariefsam/telemarketing/usecase"
	"github.com/ariefsam/telemarketing/usecase/dependency/mockdependency"
	"github.com/stretchr/testify/assert"
)

func TestCreateCustomer(t *testing.T) {
	var u usecase.Usecase
	var mockCustomerRepository mockdependency.CustomerRepository
	u.CustomerRepository = &mockCustomerRepository
	customer := entity.Customer{
		ID:          "generated",
		PhoneNumber: "082131",
		DataSource:  "Ovo",
	}
	filter := entity.FilterCustomer{
		PhoneNumber: &customer.PhoneNumber,
	}
	mockCustomerRepository.On("Get", filter, 1).Return([]entity.Customer{}, nil)
	mockCustomerRepository.On("GenerateID").Return("generated")
	mockCustomerRepository.On("Save", customer).Return(nil)
	err := u.CreateCustomer(customer)
	assert.NoError(t, err)
	mockCustomerRepository.AssertCalled(t, "Save", customer)

}

func TestCreateCustomerPhoneExist(t *testing.T) {
	var u usecase.Usecase
	var mockCustomerRepository mockdependency.CustomerRepository
	u.CustomerRepository = &mockCustomerRepository
	customer := entity.Customer{
		ID:          "generated",
		PhoneNumber: "082131",
		DataSource:  "Ovo",
	}
	filter := entity.FilterCustomer{
		PhoneNumber: &customer.PhoneNumber,
	}
	mockCustomerRepository.On("Get", filter, 1).Return([]entity.Customer{customer}, nil)
	mockCustomerRepository.On("GenerateID").Return("generated")
	mockCustomerRepository.On("Save", customer).Return(nil)
	err := u.CreateCustomer(customer)
	assert.Error(t, err)
	mockCustomerRepository.AssertNotCalled(t, "Save", customer)

}

func TestCreateCustomerNoSource(t *testing.T) {
	var u usecase.Usecase
	var mockCustomerRepository mockdependency.CustomerRepository
	u.CustomerRepository = &mockCustomerRepository
	customer := entity.Customer{
		ID:          "generated",
		PhoneNumber: "082131",
	}
	filter := entity.FilterCustomer{
		PhoneNumber: &customer.PhoneNumber,
	}
	mockCustomerRepository.On("Get", filter, 1).Return([]entity.Customer{}, nil)
	mockCustomerRepository.On("GenerateID").Return("generated")
	mockCustomerRepository.On("Save", customer).Return(nil)
	err := u.CreateCustomer(customer)
	assert.Error(t, err)
	mockCustomerRepository.AssertNotCalled(t, "Save", customer)

}
