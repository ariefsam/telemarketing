package restapi

import (
	"errors"
	"net/http"

	"time"

	"github.com/ariefsam/telemarketing/entity"
)

func (api *RestAPI) AssignCustomer(w http.ResponseWriter, r *http.Request) {
	post := getPostModel(r)

	telemarketer, err := api.authOrResponseError(post, w)
	if err != nil {
		return
	}
	emptyString := ""
	filter := entity.FilterCustomer{
		TelemarketerID: &emptyString,
	}
	customers, err := api.Usecase.GetCustomer(filter, 1)
	if err != nil {
		responseError(w, err, http.StatusBadGateway)
		return
	}
	if len(customers) == 0 {
		responseError(w, errors.New("No customer left"), http.StatusBadGateway)
		return
	}
	customer := customers[0]

	response := map[string]interface{}{}
	err = api.Usecase.ValidateAssignCustomer(telemarketer.ID, customer.ID)
	if err != nil {
		responseError(w, err, http.StatusBadGateway)
		return
	}
	err = api.Usecase.AssignCustomer(telemarketer.ID, customer.ID)
	if err != nil {
		response["Error"] = err.Error()
		JSONView(w, response, http.StatusBadGateway)
		return
	}
	time.Sleep(300 * time.Millisecond)
	response["Customer"] = customer
	JSONView(w, response, http.StatusOK)
}
