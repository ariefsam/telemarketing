package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariefsam/telemarketing/entity"
	"github.com/ariefsam/telemarketing/usecase"
	"github.com/ariefsam/telemarketing/usecase/dependency/mockdependency"
)

func TestLoginByFirebase(t *testing.T) {
	var u usecase.Usecase
	var mockFirebaseAuth mockdependency.FirebaseAuth
	u.FirebaseAuth = &mockFirebaseAuth

	firebaseToken := "dummyfirebasetoken"
	expectedToken := "tokengeneratedbysystem"
	expectedValid := true

	expectedAuthData := entity.AuthData{
		Email:   "aa@gmail.com",
		Picture: "picture.com/x.png",
	}

	mockFirebaseAuth.On("ValidateIDToken", firebaseToken).Return(expectedAuthData, expectedValid, nil)

	var mockTelemarketer mockdependency.TelemarketerRepository
	u.TelemarketerRepository = &mockTelemarketer

	expectedTelemarketer := entity.Telemarketer{
		Email: expectedAuthData.Email,
		Name:  "Dadang",
	}
	filterTelemarketing := entity.FilterTelemarketer{
		Email: expectedAuthData.Email,
	}
	mockTelemarketer.On("Get", filterTelemarketing, 1).Return([]entity.Telemarketer{expectedTelemarketer}, nil)

	var mockTokenService mockdependency.TokenService
	u.TokenService = &mockTokenService

	mockTokenService.On("Create", expectedTelemarketer).Return(expectedToken)

	token, telemarketer, isValid, err := u.LoginByFirebase(firebaseToken)
	assert.NoError(t, err)
	assert.Equal(t, expectedToken, token)
	assert.Equal(t, expectedTelemarketer, telemarketer)
	assert.Equal(t, expectedValid, isValid)

}
