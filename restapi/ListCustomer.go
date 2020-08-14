package restapi

import (
	"encoding/json"
	"net/http"
)

func (api *RestAPI) ListCustomer(w http.ResponseWriter, r *http.Request) {
	var post RequestModel
	json.NewDecoder(r.Body).Decode(&post)

	telemarketer, err := api.authOrResponseError(post, w)
	if err != nil {
		return
	}

	response := map[string]interface{}{}

	if post.FilterCustomer.TelemarketerEmail != nil && telemarketer.IsAdmin == false {
		response["Error"] = "FilterCustomer TelemarketerEmail forbidden"
		JSONView(w, response, http.StatusForbidden)
		return

	}

	if telemarketer.IsAdmin == false {
		post.FilterCustomer.TelemarketerEmail = &telemarketer.Email
	}

	customers, err := api.Usecase.GetCustomer(post.FilterCustomer, post.Limit)
	if err != nil {
		response["Error"] = err.Error()
		JSONView(w, response, http.StatusBadGateway)
		return
	}
	response["Customers"] = customers
	JSONView(w, response, http.StatusOK)
	return
}
