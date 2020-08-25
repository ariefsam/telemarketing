package restapi

import (
	"net/http"
)

func (api *RestAPI) ListCustomer(w http.ResponseWriter, r *http.Request) {
	post := getPostModel(r)

	telemarketer, err := api.authOrResponseError(post, w)
	if err != nil {
		return
	}

	response := map[string]interface{}{}

	if post.FilterCustomer.TelemarketerID != nil && telemarketer.IsAdmin == false {
		response["Error"] = "FilterCustomer TelemarketerID forbidden"
		JSONView(w, response, http.StatusForbidden)
		return

	}

	if telemarketer.IsAdmin == false {
		post.FilterCustomer.TelemarketerID = &telemarketer.Email
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
