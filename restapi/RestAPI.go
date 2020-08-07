package restapi

import "github.com/ariefsam/telemarketing/restapi/usecaseinterface"

type RestAPI struct {
	Usecase usecaseinterface.Usecase
}
