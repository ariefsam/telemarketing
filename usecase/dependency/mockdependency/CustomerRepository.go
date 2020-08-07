package mockdependency

import (
	"github.com/ariefsam/telemarketing/entity"
	"github.com/stretchr/testify/mock"
)

type CustomerRepository struct {
	mock.Mock
}

func (m *CustomerRepository) Save(customer entity.Customer) (err error) {
	args := m.Called(customer)
	err = args.Error(0)
	return

}
func (m *CustomerRepository) Get(filter entity.FilterCustomer, limit int) (customers []entity.Customer, err error) {
	args := m.Called(filter, limit)
	customers = args.Get(0).([]entity.Customer)
	err = args.Error(1)
	return
}
