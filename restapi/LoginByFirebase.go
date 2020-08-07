package restapi

import (
	"encoding/json"
	"net/http"
)

func (api *RestAPI) LoginByFirebase(w http.ResponseWriter, r *http.Request) {
	var post map[string]string
	var firebaseToken string
	var ok bool
	response := map[string]interface{}{}
	json.NewDecoder(r.Body).Decode(&post)

	if firebaseToken, ok = post["FirebaseToken"]; !ok {
		response["Error"] = "Parameter 'firebaseToken' is needed"
		JSONView(w, response, http.StatusBadRequest)
		return
	}

	token, telemarketer, isValid, err := api.Usecase.LoginByFirebase(firebaseToken)
	if err != nil {
		response["Error"] = "Login failed"
		JSONView(w, response, http.StatusBadGateway)
		return
	}
	if !isValid {
		response["Error"] = "Invalid Firebase Token"
		JSONView(w, response, http.StatusBadGateway)
		return
	}
	response["Token"] = token
	response["Telemarketer"] = telemarketer
	JSONView(w, response, http.StatusOK)
}
