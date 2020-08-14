package restapi_test

import (
	"github.com/ariefsam/telemarketing/entity"
	"github.com/ariefsam/telemarketing/restapi"
	"github.com/ariefsam/telemarketing/restapi/usecaseinterface/mockusecase"
)

func setupMockParseToken(mockUsecase *mockusecase.Usecase) (dummyToken string, expectedTelemarketer entity.Telemarketer) {
	dummyToken = "dTOkne"
	expectedTelemarketer = dummyTelemarketerNotAdmin()
	mockUsecase.On("ParseToken", dummyToken).Return(true, expectedTelemarketer)
	return
}

func setupAPIAndUsecase() (api *restapi.RestAPI, mockUsecase *mockusecase.Usecase) {
	var m mockusecase.Usecase
	api = new(restapi.RestAPI)
	mockUsecase = &m
	api.Usecase = mockUsecase
	return
}
