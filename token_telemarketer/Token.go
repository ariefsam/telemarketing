package token_telemarketer

import (
	"log"
	"time"

	"github.com/ariefsam/telemarketing/entity"
	"github.com/dgrijalva/jwt-go"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

type Token struct {
	Secret []byte
}
type MyCustomClaims struct {
	Telemarketer entity.Telemarketer
	jwt.StandardClaims
}

func (t *Token) Create(telemarketer entity.Telemarketer) (tokenString string) {

	claims := MyCustomClaims{
		telemarketer,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 604800,
			Issuer:    "go-chat",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ = token.SignedString(t.Secret)
	return
}
func (t *Token) Parse(tokenString string) (isValid bool, telemarketer entity.Telemarketer) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return t.Secret, nil
	})

	if err != nil {
		log.Println(err)
		return
	}

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		isValid = true
		telemarketer = claims.Telemarketer
	} else {

	}
	return

}
