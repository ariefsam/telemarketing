package usecase_test

import (
	"testing"

	"github.com/ariefsam/telemarketing/entity"
	"github.com/ariefsam/telemarketing/usecase"
	"github.com/ariefsam/telemarketing/usecase/dependency/mockdependency"
	"github.com/stretchr/testify/assert"
)

func TestParseToken(t *testing.T) {
	var u usecase.Usecase
	var mockTokenService mockdependency.TokenService
	u.TokenService = &mockTokenService
	token := "dummyToken"
	telemarketer := entity.Telemarketer{
		Name:  "Ad",
		Email: "email@gmail.com",
	}
	valid := true
	mockTokenService.On("Parse", token).Return(valid, telemarketer)
	isValid, getTelemarketer := u.ParseToken(token)
	assert.Equal(t, valid, isValid)
	assert.Equal(t, telemarketer, getTelemarketer)
}
