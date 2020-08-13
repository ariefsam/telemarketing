package restapi

import "github.com/ariefsam/telemarketing/entity"

type RequestModel struct {
	Token          *string
	Limit          int
	FilterCustomer entity.FilterCustomer
}
