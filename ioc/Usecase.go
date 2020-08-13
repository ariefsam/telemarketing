package ioc

import (
	"github.com/ariefsam/telemarketing/configuration"
	"github.com/ariefsam/telemarketing/lib/firebaseauth"
	"github.com/ariefsam/telemarketing/lib/repository"
	"github.com/ariefsam/telemarketing/lib/timer"
	"github.com/ariefsam/telemarketing/lib/token_telemarketer"
	"github.com/ariefsam/telemarketing/usecase"
)

func Usecase() (usecase usecase.Usecase) {
	repository.FirebaseAccountPath = configuration.FirebaseCredentialPath
	usecase.CustomerRepository = &repository.Customer{}
	usecase.CallLogRepository = &repository.CallLog{}
	usecase.TokenService = &token_telemarketer.Token{}
	usecase.FirebaseAuth = &firebaseauth.FirebaseAuth{
		CredentialPath: configuration.FirebaseCredentialPath,
	}
	usecase.TelemarketerRepository = &repository.Telemarketer{}
	usecase.Timer = &timer.Timer{}
	return
}
