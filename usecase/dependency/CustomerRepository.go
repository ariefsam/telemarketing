package dependency

import "github.com/ariefsam/telemarketing/entity"

type CustomerRepository interface {
	Save(customer entity.Customer) (err error)
	Get(filter entity.FilterCustomer, limit int) (customers []entity.Customer, err error)
}
