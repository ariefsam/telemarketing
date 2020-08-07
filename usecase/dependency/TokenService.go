package dependency

import "github.com/ariefsam/telemarketing/entity"

type TokenService interface {
	Create(telemarketer entity.Telemarketer) (token string)
	Parse(token string) (isValid bool, user entity.Telemarketer)
}
