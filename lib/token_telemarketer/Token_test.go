package token_telemarketer_test

import (
	"testing"

	"github.com/ariefsam/telemarketing/entity"
	"github.com/ariefsam/telemarketing/lib/token_telemarketer"
	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
	var telemarketer entity.Telemarketer
	telemarketer.Name = "name telemarketer"
	var tokenService token_telemarketer.Token
	tokenService.Secret = []byte("xxx")
	token := tokenService.Create(telemarketer)
	assert.NotEmpty(t, token)
	isValid, parseTelemarketer := tokenService.Parse(token)
	assert.True(t, isValid)
	assert.Equal(t, telemarketer, parseTelemarketer)
}
