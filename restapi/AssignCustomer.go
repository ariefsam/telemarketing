package restapi

import (
	"encoding/json"
	"net/http"
)

func (api *RestAPI) AssignCustomer(w http.ResponseWriter, r *http.Request) {
	var post RequestModel
	json.NewDecoder(r.Body).Decode(&post)

	telemarketer, err := api.authOrResponseError(post, w)
	if err != nil {
		return
	}

	response := map[string]interface{}{}
	customer, err := api.Usecase.AssignCustomer(telemarketer.Email)
	if err != nil {
		response["Error"] = err.Error()
		JSONView(w, response, http.StatusBadGateway)
		return
	}

	response["Customer"] = customer
	JSONView(w, response, http.StatusOK)
}
