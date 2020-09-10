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

	telemarketer, isValid, err = u.getTelemarketerByEmail(authData.Email)
	if err != nil {
		return
	}

	token = u.TokenService.Create(telemarketer)
	return
}
