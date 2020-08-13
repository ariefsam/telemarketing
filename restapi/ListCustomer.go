package restapi

import (
	"encoding/json"
	"net/http"

	"github.com/ariefsam/telemarketing/entity"
)

func (api *RestAPI) ListCustomer(w http.ResponseWriter, r *http.Request) {
	var post RequestModel
	response := map[string]interface{}{}
	json.NewDecoder(r.Body).Decode(&post)

	telemarketer, err := api.authOrResponseError(post, w)
	if err != nil {
		return
	}
	limit := api.parseInteger(post.Limit)
	filter := entity.FilterCustomer{
		TelemarketerID: &telemarketer.Email,
	}
	customers, err := api.Usecase.GetCustomer(filter, limit)
	if err != nil {
		response["Error"] = err.Error()
		JSONView(w, response, http.StatusBadGateway)
		return
	}
	response["Customers"] = customers
	JSONView(w, response, http.StatusOK)
	return
}
