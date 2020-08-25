package restapi_test

import (
	"testing"

	"github.com/ariefsam/telemarketing/entity"
)

func TestCreateTelemarketerNotAdmin(t *testing.T) {
	api, mockUsecase := setupAPIAndUsecase()
	dummyToken, expectedTelemarketer := setupMockParseToken(mockUsecase)
	expectedTelemarketer.IsAdmin = false

	request := map[string]interface{}{
		"Token": dummyToken,
		"Telemarketer": map[string]interface{}{
			"Name":    "Arief",
			"Email":   "arief@fsn.co.id",
			"IsAdmin": false,
		},
	}

	expectedResponse := map[string]interface{}{
		"Error": "Unauthorized",
	}
	assertRequestResponse(t, api.SaveTelemarketer, request, expectedResponse)
}

func TestCreateTelemarketerAdmin(t *testing.T) {
	api, mockUsecase := setupAPIAndUsecase()
	dummyToken, _ := setupMockParseTokenAdmin(mockUsecase)

	request := map[string]interface{}{
		"Token": dummyToken,
		"Telemarketer": map[string]interface{}{
			"Name":    "Arief",
			"Email":   "arief@fsn.co.id",
			"IsAdmin": false,
		},
	}

	savedTelemarketer := entity.Telemarketer{
		Name:    "Arief",
		Email:   "arief@fsn.co.id",
		IsAdmin: false,
	}

	mockUsecase.On("CreateTelemarketer", savedTelemarketer).Return(nil)
	expectedResponse := map[string]interface{}{
		"Status": "ok",
	}
	assertRequestResponse(t, api.CreateTelemarketer, request, expectedResponse)
	mockUsecase.AssertCalled(t, "CreateTelemarketer", savedTelemarketer)
}
