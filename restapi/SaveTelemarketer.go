package restapi

import "net/http"

func (api *RestAPI) SaveTelemarketer(w http.ResponseWriter, r *http.Request) {
	post := getPostModel(r)

	_, err := api.authAdminOrResponseError(post, w)
	if err != nil {
		return
	}

	err = api.Usecase.SaveTelemarketer(post.Telemarketer)
	if err != nil {
		responseError(w, err, http.StatusBadGateway)
		return
	}

	response := map[string]interface{}{
		"Status": "ok",
	}
	JSONView(w, response, http.StatusOK)
	return

}
