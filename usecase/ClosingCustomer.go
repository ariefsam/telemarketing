package usecase

import (
	"log"

	"github.com/ariefsam/telemarketing/entity"
)

func (u *Usecase) ClosingCustomer(customer entity.Customer) (err error) {

	err = u.ValidateClosingCustomer(customer)
	if err != nil {
		return
	}
	customer.ClosingTimestamp = u.Timer.CurrentTimestamp()
	customer.IsClosing = true
	customer.TimestampUpdated = u.Timer.CurrentTimestamp()

	err = u.CustomerRepository.Save(customer)
	return
}

func (u *Usecase) ValidateClosingCustomer(customer entity.Customer) (err error) {
	var currentCustomer entity.Customer

	filter := entity.FilterCustomer{
		ID: &customer.ID,
	}
	temp, err := u.CustomerRepository.Get(filter, 1)
	if err != nil {
		return err
	}
	if len(temp) == 1 {
		currentCustomer = temp[0]
	}

	if currentCustomer.IsClosing {
		log.Println("customer ", customer.ID, " ", customer.Name, " already closed")
		return
	}
	return
}
