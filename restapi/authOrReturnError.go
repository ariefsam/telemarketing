package restapi

import (
	"errors"
	"net/http"

	"github.com/ariefsam/telemarketing/entity"
)

func (api *RestAPI) authOrResponseError(post RequestModel, w http.ResponseWriter) (telemarketer entity.Telemarketer, err error) {
	response := map[string]string{}
	if post.Token == nil {
		err = errors.New("Token is nil")
		response["Error"] = "Token is needed"
		JSONView(w, response, http.StatusUnauthorized)
		return
	}
	if *post.Token == "" {
		err = errors.New("Token is empty")
		response["Error"] = "Token is needed"
		JSONView(w, response, http.StatusUnauthorized)
	}
	var isValid bool
	isValid, telemarketer = api.Usecase.ParseToken(*post.Token)
	if !isValid {
		err = errors.New("Invalid Token")
		response["Error"] = err.Error()
		JSONView(w, response, http.StatusUnauthorized)
	}
	return
}

func (api *RestAPI) authAdminOrResponseError(post RequestModel, w http.ResponseWriter) (telemarketer entity.Telemarketer, err error) {
	response := map[string]string{}
	telemarketer, err = api.authOrResponseError(post, w)
	if err != nil {
		return
	}
	if !telemarketer.IsAdmin {
		err = errors.New("Unauthorized")
		response["Error"] = err.Error()
		JSONView(w, response, http.StatusUnauthorized)
		return
	}
	return
}
