package restapi_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertRequestResponse(t *testing.T, handlerFunc http.HandlerFunc, requestBody map[string]interface{}, expectedResponse map[string]interface{}) {
	bodyByte, _ := json.Marshal(requestBody)
	bodyReader := bytes.NewReader(bodyByte)
	req := httptest.NewRequest("POST", "http://localhost/", bodyReader)
	w := httptest.NewRecorder()

	expectedBodyResponse, _ := json.Marshal(expectedResponse)

	handlerFunc(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	actualResponse := map[string]interface{}{}
	json.Unmarshal(body, &actualResponse)
	assert.Equal(t, expectedBodyResponse, body)
}
