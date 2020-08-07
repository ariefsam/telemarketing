package usecase

import "github.com/ariefsam/telemarketing/usecase/dependency"

type Usecase struct {
	CustomerRepository     dependency.CustomerRepository
	CallLogRepository      dependency.CallLogRepository
	TokenService           dependency.TokenService
	FirebaseAuth           dependency.FirebaseAuth
	TelemarketerRepository dependency.TelemarketerRepository
}
