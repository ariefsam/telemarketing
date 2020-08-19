package restapi

import (
	"log"
	"net/http"
)

func responseError(w http.ResponseWriter, err error, statusCode int) {
	response := map[string]interface{}{}
	if err != nil {
		log.Println(err)
		response["Error"] = err.Error()
		JSONView(w, response, statusCode)
		return
	}
}
