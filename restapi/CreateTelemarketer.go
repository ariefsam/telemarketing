package restapi

import (
	"net/http"
	"time"
)

func (api *RestAPI) CreateTelemarketer(w http.ResponseWriter, r *http.Request) {
	post := getPostModel(r)

	_, err := api.authAdminOrResponseError(post, w)
	if err != nil {
		return
	}
	post.Telemarketer.ID = api.Usecase.GenerateID()
	err = api.Usecase.CreateTelemarketer(post.Telemarketer)
	if err != nil {
		responseError(w, err, http.StatusBadGateway)
		return
	}

	response := map[string]interface{}{
		"Status": "ok",
	}
	time.Sleep(300 * time.Millisecond)
	JSONView(w, response, http.StatusOK)
	return

}
