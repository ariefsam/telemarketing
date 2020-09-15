package usecase

import (
	"errors"

	"github.com/ariefsam/telemarketing/entity"
)

func (u *Usecase) AssignCustomer(telemarketerID string, customerID string) (err error) {
	filter := entity.FilterCustomer{
		TelemarketerID: new(string),
	}
	err = u.ValidateAssignCustomer(telemarketerID, customerID)
	if err != nil {
		return
	}
	customers, err := u.CustomerRepository.Get(filter, 1)
	if err != nil {
		return
	}
	if len(customers) == 0 {
		err = errors.New("No customer left")
		return
	}
	customer := customers[0]
	customer.TelemarketerID = telemarketerID
	err = u.CustomerRepository.Save(customer)
	return
}

func (u *Usecase) ValidateAssignCustomer(telemarketerID string, customerID string) (err error) {

	return
}
