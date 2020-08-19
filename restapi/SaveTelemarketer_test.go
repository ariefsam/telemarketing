package restapi_test

import (
	"testing"

	"github.com/ariefsam/telemarketing/entity"
)

func TestSaveTelemarketerNotAdmin(t *testing.T) {
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

func TestSaveTelemarketerAdmin(t *testing.T) {
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

	mockUsecase.On("SaveTelemarketer", savedTelemarketer).Return(nil)
	expectedResponse := map[string]interface{}{
		"Telemarketer": map[string]interface{}{
			"Name":    "Arief",
			"Email":   "arief@fsn.co.id",
			"IsAdmin": false,
		},
	}
	assertRequestResponse(t, api.SaveTelemarketer, request, expectedResponse)
	mockUsecase.AssertCalled(t, "SaveTelemarketer", savedTelemarketer)
}
