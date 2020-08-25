package usecase

import (
	"errors"

	"github.com/ariefsam/telemarketing/entity"
)

func (u *Usecase) AssignCustomer(telemarketerID string) (customer entity.Customer, err error) {
	filter := entity.FilterCustomer{
		TelemarketerID: new(string),
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
	customer.TelemarketerID = telemarketerID
	err = u.CustomerRepository.Save(customer)
	return
}
