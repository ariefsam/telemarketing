package usecase

import (
	"errors"

	"github.com/ariefsam/telemarketing/entity"
)

func (u *Usecase) CreateCustomer(customer entity.Customer) (err error) {
	err = u.ValidateCreateCustomer(customer)
	if err != nil {
		return
	}
	err = u.CustomerRepository.Save(customer)
	return
}

func (u *Usecase) ValidateCreateCustomer(customer entity.Customer) (err error) {
	if customer.DataSource == "" {
		err = errors.New("DataSource needed")
		return
	}
	if customer.ID == "" {
		err = errors.New("ID needed")
		return
	}
	filter := entity.FilterCustomer{
		PhoneNumber: &customer.PhoneNumber,
	}
	customers, err := u.CustomerRepository.Get(filter, 1)
	if err != nil {
		return
	}
	if len(customers) > 0 {
		err = errors.New("Phone number already used by other customer")
	}

	return
}
