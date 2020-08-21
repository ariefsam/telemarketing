package restapi_test

import (
	"testing"

	"github.com/ariefsam/telemarketing/entity"
)

func TestGetTelemarketerAdmin(t *testing.T) {
	api, mockUsecase := setupAPIAndUsecase()
	dummyToken, _ := setupMockParseTokenAdmin(mockUsecase)
	request := map[string]interface{}{
		"Token": dummyToken,
		"Limit": 1,
	}
	filter := entity.FilterTelemarketer{}
	limit := 1
	telemarketers := []entity.Telemarketer{entity.Telemarketer{Email: "x@gmail.com"}}
	mockUsecase.On("GetTelemarketer", filter, limit).Return(telemarketers, nil)

	response := map[string]interface{}{
		"Telemarketers": telemarketers,
	}
	assertRequestResponse(t, api.GetTelemarketer, request, response)
}

func TestGetTelemarketerNotAdmin(t *testing.T) {
	api, mockUsecase := setupAPIAndUsecase()
	dummyToken, _ := setupMockParseToken(mockUsecase)
	request := map[string]interface{}{
		"Token": dummyToken,
		"Limit": 1,
	}
	filter := entity.FilterTelemarketer{}
	limit := 1
	telemarketers := []entity.Telemarketer{entity.Telemarketer{Email: "x@gmail.com"}}
	mockUsecase.On("GetTelemarketer", filter, limit).Return(telemarketers, nil)

	response := map[string]interface{}{
		"Error": "Unauthorized",
	}
	assertRequestResponse(t, api.GetTelemarketer, request, response)
}
