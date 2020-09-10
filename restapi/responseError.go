package restapi

import (
	"net/http"
)

func responseError(w http.ResponseWriter, err error, statusCode int) {
	response := map[string]interface{}{}
	if err != nil {
		response["Error"] = err.Error()
		JSONView(w, response, statusCode)
		return
	}
}
