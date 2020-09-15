package usecase

import "github.com/ariefsam/telemarketing/entity"

func (u *Usecase) Call(telemarketer entity.Telemarketer, customer entity.Customer, status string, timestamp int64) (err error) {

	return
}

func (u *Usecase) ValidateCall() {
	return
}
