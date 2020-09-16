package usecase

import "github.com/ariefsam/telemarketing/entity"

func (u *Usecase) SaveCustomer(customer entity.Customer) (err error) {
	var checkCurrentCustomer bool
	var currentCustomer entity.Customer
	if customer.IsClosing == true {
		checkCurrentCustomer = true
	}

	if checkCurrentCustomer {
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
	}

	if customer.IsClosing {
		if !currentCustomer.IsClosing {
			customer.ClosingTimestamp = u.Timer.CurrentTimestamp()
		}
	}
	err = u.CustomerRepository.Save(customer)
	return
}
