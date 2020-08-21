package usecase

import "github.com/ariefsam/telemarketing/entity"

func (u *Usecase) GetTelemarketer(filter entity.FilterTelemarketer, limit int) (telemarketers []entity.Telemarketer, err error) {
	telemarketers, err = u.TelemarketerRepository.Get(filter, limit)
	return
}
