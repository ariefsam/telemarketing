package usecase

import "github.com/ariefsam/telemarketing/entity"

func (u *Usecase) LoginByFirebase(firebaseToken string) (token string, telemarketer entity.Telemarketer, isValid bool, err error) {
	var authData entity.AuthData
	authData, isValid, err = u.FirebaseAuth.ValidateIDToken(firebaseToken)
	if err != nil {
		return
	}
	if !isValid {
		return
	}

	var telemarketers []entity.Telemarketer
	filter := entity.FilterTelemarketer{
		Email: authData.Email,
	}
	telemarketers, err = u.TelemarketerRepository.Get(filter, 1)
	if err != nil {
		return
	}
	if len(telemarketers) == 0 {
		isValid = false
		return
	}

	telemarketer = telemarketers[0]
	token = u.TokenService.Create(telemarketer)
	return
}
