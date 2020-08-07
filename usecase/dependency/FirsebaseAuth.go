package dependency

import "github.com/ariefsam/telemarketing/entity"

type FirebaseAuth interface {
	ValidateIDToken(token string) (authData entity.AuthData, isValid bool, err error)
}
