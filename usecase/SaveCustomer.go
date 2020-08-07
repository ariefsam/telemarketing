package usecase

import "github.com/ariefsam/telemarketing/entity"

func (u *Usecase) SaveCustomer(customer entity.Customer) (err error) {
	err = u.CustomerRepository.Save(customer)
	return
}
