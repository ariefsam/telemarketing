package restapi

import "github.com/ariefsam/telemarketing/entity"

type RequestModel struct {
	Token              *string
	Limit              int
	FilterCustomer     entity.FilterCustomer
	PhoneNumber        string
	Status             string
	Telemarketer       entity.Telemarketer
	FilterTelemarketer entity.FilterTelemarketer
	FilterCallLog      entity.FilterCallLog
	Customer           *entity.Customer
}
