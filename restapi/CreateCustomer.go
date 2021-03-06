package restapi

import (
	"errors"
	"net/http"
)

func (api *RestAPI) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	post := getPostModel(r)

	telemarketer, err := api.authOrResponseError(post, w)
	if err != nil {
		return
	}
	if post.Customer == nil {
		responseError(w, errors.New("No Customer"), http.StatusBadGateway)
		return
	}
	// post.Telemarketer.ID = api.Usecase.GenerateID()
	post.Customer.ID = api.Usecase.GenerateID()
	post.Customer.CreatedBy = telemarketer.ID
	post.Customer.TelemarketerID = telemarketer.ID
	err = api.Usecase.CreateCustomer(*post.Customer)
	if err != nil {
		responseError(w, err, http.StatusBadGateway)
		return
	}

	response := map[string]interface{}{
		"Status": "ok",
	}

	JSONView(w, response, http.StatusOK)
	return

}
