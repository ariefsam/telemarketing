package restapi

import (
	"net/http"

	"github.com/ariefsam/telemarketing/entity"
)

func (api *RestAPI) Me(w http.ResponseWriter, r *http.Request) {
	post := getPostModel(r)
	telemarketer, err := api.authOrResponseError(post, w)
	if err != nil {
		return
	}
	filter := entity.FilterTelemarketer{
		ID: &telemarketer.ID,
	}

	telemarketers, err := api.Usecase.GetTelemarketer(filter, post.Limit)
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
