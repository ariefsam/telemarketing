package usecase

import "github.com/ariefsam/telemarketing/entity"

func (u *Usecase) getTelemarketerByEmail(email string) (telemarketer entity.Telemarketer, isExist bool, err error) {
	filter := entity.FilterTelemarketer{
		Email: &email,
	}
	var telemarketers []entity.Telemarketer
	telemarketers, err = u.TelemarketerRepository.Get(filter, 1)
	if err != nil {
		return
	}
	if len(telemarketers) > 0 {
		isExist = true
		telemarketer = telemarketers[0]
	}
	return
}
