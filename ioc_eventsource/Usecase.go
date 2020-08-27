package ioc_eventsource

import (
	"github.com/ariefsam/telemarketing/configuration"
	"github.com/ariefsam/telemarketing/ioc"
	"github.com/ariefsam/telemarketing/usecase_eventsource"
)

func Usecase() (usecase usecase_eventsource.Usecase) {
	usecase_eventsource.FirebaseAccountPath = configuration.FirebaseCredentialPath
	usecase.Usecase = ioc.Usecase()
	return
}
