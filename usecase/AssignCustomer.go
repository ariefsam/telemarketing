package usecase

import (
	"errors"

	"github.com/ariefsam/telemarketing/entity"
)

func (u *Usecase) AssignCustomer(telemarketingID string) (customer entity.Customer, err error) {
	filter := entity.FilterCustomer{
		TelemarketerID: "Not Set",
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
	customer.TelemarketerID = telemarketingID
	err = u.CustomerRepository.Save(customer)
	return
}
