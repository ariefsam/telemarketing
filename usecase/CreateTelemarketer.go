package usecase

import (
	"errors"

	"github.com/ariefsam/telemarketing/entity"
)

func (u *Usecase) CreateTelemarketer(telemarketer entity.Telemarketer) (err error) {
	err = u.ValidateCreateTelemarketer(telemarketer)
	if err != nil {
		return
	}
	err = u.TelemarketerRepository.Save(telemarketer)
	return
}

func (u *Usecase) ValidateCreateTelemarketer(telemarketer entity.Telemarketer) (err error) {
	var isExist bool
	_, isExist, err = u.getTelemarketerByEmail(telemarketer.Email)
	if err != nil {
		return
	}
	if isExist {
		err = errors.New("Email already exist")
		return
	}
	return
}
