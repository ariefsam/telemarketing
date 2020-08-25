package restapi

import (
	"net/http"
)

func (api *RestAPI) GetCallLog(w http.ResponseWriter, r *http.Request) {
	post := getPostModel(r)
	telemarketer, err := api.authOrResponseError(post, w)
	if err != nil {
		return
	}
	filter := post.FilterCallLog
	if !telemarketer.IsAdmin {
		filter.TelemarketerID = &telemarketer.ID
	}

	CallLogs, err := api.Usecase.GetCallLog(filter, post.Limit)
	if err != nil {
		responseError(w, err, http.StatusBadGateway)
	}
	response := map[string]interface{}{
		"CallLogs": CallLogs,
	}
	JSONView(w, response, http.StatusOK)
}
