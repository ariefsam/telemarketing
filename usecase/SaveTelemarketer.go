package usecase

import "github.com/ariefsam/telemarketing/entity"

func (u *Usecase) SaveTelemarketer(telemarketer entity.Telemarketer) (err error) {
	err = u.TelemarketerRepository.Save(telemarketer)
	return
}
