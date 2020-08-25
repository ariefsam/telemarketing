package restapi

import (
	"errors"
	"net/http"

	"github.com/ariefsam/telemarketing/entity"
)

func (api *RestAPI) Call(w http.ResponseWriter, r *http.Request) {
	post := getPostModel(r)

	telemarketer, err := api.authOrResponseError(post, w)
	if err != nil {
		return
	}

	filter := entity.FilterCustomer{
		PhoneNumber: &post.PhoneNumber,
	}

	customers, err := api.Usecase.GetCustomer(filter, 1)
	if err != nil {
		responseError(w, err, http.StatusBadGateway)
		return
	}
	if len(customers) == 0 {
		responseError(w, errors.New("Customer Not Found"), http.StatusBadRequest)
		return
	}
	customer := customers[0]

	if telemarketer.Email != customers[0].TelemarketerID {
		responseError(w, errors.New("Forbidden Customer"), http.StatusBadRequest)
		return
	}

	status := post.Status
	timestamp := api.Usecase.CurrentTimestamp()
	callLog := entity.CallLog{
		Name:           customer.Name,
		PhoneNumber:    customer.PhoneNumber,
		Status:         status,
		Timestamp:      timestamp,
		TelemarketerID: telemarketer.ID,
	}
	err = api.Usecase.SaveCallLog(callLog)
	if err != nil {
		responseError(w, err, http.StatusBadGateway)
		return
	}

	customer.Status = status
	customer.TelemarketerID = telemarketer.ID
	customer.TimestampUpdated = timestamp
	err = api.Usecase.SaveCustomer(customer)
	if err != nil {
		responseError(w, err, http.StatusBadGateway)
		return
	}
	response := map[string]interface{}{
		"CallLog":  callLog,
		"Customer": customer,
	}
	JSONView(w, response, http.StatusOK)
	return
}
