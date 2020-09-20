package usecase

import (
	"errors"

	"github.com/ariefsam/telemarketing/lib/timer"

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
	if err != nil {
		return
	}
	myTime := timer.Unix(0, u.Timer.CurrentTimestamp()).In(loc)
	u.ReportRepository.Increment(customer.TelemarketerID, "CLOSING", 1, myTime)
	u.ReportRepository.Increment(customer.TelemarketerID, "BUYAMOUNT", customer.BuyAmount, myTime)
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
		err = errors.New("customer " + customer.ID + " " + customer.Name + " already closed")
		return
	}
	if customer.BuyAmount == 0 {
		err = errors.New("buy amount " + customer.ID + " " + customer.Name + " must be set")
		return
	}
	return
}
