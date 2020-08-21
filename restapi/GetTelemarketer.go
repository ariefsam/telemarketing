package restapi

import (
	"net/http"
)

func (api *RestAPI) GetTelemarketer(w http.ResponseWriter, r *http.Request) {
	post := getPostModel(r)
	_, err := api.authAdminOrResponseError(post, w)
	if err != nil {
		return
	}

	telemarketers, err := api.Usecase.GetTelemarketer(post.FilterTelemarketer, post.Limit)
	if err != nil {
		responseError(w, err, http.StatusBadGateway)
		return
	}
	response := map[string]interface{}{
		"Telemarketers": telemarketers,
	}
	JSONView(w, response, http.StatusOK)
	return
}
