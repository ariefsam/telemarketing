package restapi_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/ariefsam/telemarketing/entity"
	"github.com/stretchr/testify/assert"
)

func TestLoginByFirebase(t *testing.T) {

	api, mockUsecase := setupAPIAndUsecase()

	firebaseToken := "firebaseTokenDummy"
	expectedTelemarketerToken := "expectedTelemarketerTokenxxx"
	expectedTelemarketer := entity.Telemarketer{
		Email: "xxx@gmail.com",
	}
	expectedIsValid := true

	expectedResponse := map[string]interface{}{
		"Token":        expectedTelemarketerToken,
		"Telemarketer": expectedTelemarketer,
	}

	expectedBodyResponse, _ := json.Marshal(expectedResponse)

	mockUsecase.On("LoginByFirebase", firebaseToken).Return(expectedTelemarketerToken, expectedTelemarketer, expectedIsValid, nil)

	bodyMap := map[string]interface{}{
		"FirebaseToken": firebaseToken,
	}
	bodyByte, _ := json.Marshal(bodyMap)
	bodyReader := bytes.NewReader(bodyByte)
	req := httptest.NewRequest("POST", "http://localhost:8885/api/login/firebase", bodyReader)
	w := httptest.NewRecorder()

	api.LoginByFirebase(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	actualResponse := map[string]interface{}{}
	json.Unmarshal(body, &actualResponse)
	assert.Equal(t, expectedResponse["Token"], actualResponse["Token"])
	assert.Equal(t, expectedBodyResponse, body)
}
