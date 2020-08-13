package usecase

import (
	"errors"

	"github.com/ariefsam/telemarketing/entity"
)

func (u *Usecase) AssignCustomer(telemarketerEmail string) (customer entity.Customer, err error) {
	filter := entity.FilterCustomer{
		TelemarketerEmail: new(string),
	}
	customers, err := u.CustomerRepository.Get(filter, 1)
	if err != nil {
		return
	}
	if len(customers) == 0 {
		err = errors.New("No customer left")
		return
	}
	customer = customers[0]
	customer.TelemarketerEmail = telemarketerEmail
	err = u.CustomerRepository.Save(customer)
	return
}
