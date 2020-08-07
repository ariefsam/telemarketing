package restapi

import (
	"errors"

	"github.com/ariefsam/telemarketing/entity"
)

func (api *RestAPI) parseToken(post map[string]interface{}) (telemarketer entity.Telemarketer, err error) {
	if _, ok := post["Token"]; !ok {
		err = errors.New("Token is needed")
		return
	}
	if post["Token"] == nil {
		err = errors.New("Token is needed")
		return
	}
	var isValid bool
	isValid, telemarketer = api.Usecase.ParseToken(post["Token"].(string))
	if !isValid {
		err = errors.New("Invalid Token")
	}
	return
}
