package usecase

import "github.com/ariefsam/telemarketing/entity"

func (u *Usecase) ParseToken(token string) (isValid bool, telemarketer entity.Telemarketer) {
	isValid, telemarketer = u.TokenService.Parse(token)
	return
}
