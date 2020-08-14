package restapi_test

import (
	"github.com/ariefsam/telemarketing/entity"
	"github.com/ariefsam/telemarketing/restapi/usecaseinterface/mockusecase"
)

func setupMockParseToken(mockUsecase *mockusecase.Usecase) (dummyToken string, expectedTelemarketer entity.Telemarketer) {
	dummyToken = "dTOkne"
	expectedTelemarketer = dummyTelemarketerNotAdmin()
	mockUsecase.On("ParseToken", dummyToken).Return(true, expectedTelemarketer)
	return
}
