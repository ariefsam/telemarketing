package usecase

import "github.com/ariefsam/telemarketing/entity"

func (u *Usecase) GetCustomer(filter entity.FilterCustomer, limit int) (customers []entity.Customer, err error) {
	customers, err = u.CustomerRepository.Get(filter, limit)
	return
}
