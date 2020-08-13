package restapi

import (
	"net/http"

	"github.com/ariefsam/telemarketing/entity"
)

func (api *RestAPI) ListCustomer(w http.ResponseWriter, r *http.Request) {

	response := map[string]interface{}{}
	customers, err := api.Usecase.GetCustomer(entity.FilterCustomer{}, 20000)
	if err != nil {
		response["Error"] = err.Error()
		JSONView(w, response, http.StatusBadGateway)
	}
	response["Customers"] = customers
	JSONView(w, response, http.StatusOK)
	return
}
