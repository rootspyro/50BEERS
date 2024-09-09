package jwt

import (
	"time"

	jwtcore "github.com/golang-jwt/jwt/v5"
	"github.com/rootspyro/50BEERS/config"
)


func Encode(email string) (string, error) {

	now := time.Now()

	token := jwtcore.NewWithClaims(jwtcore.SigningMethodHS256, jwtcore.MapClaims{
		"sub": email, // subject
		"iat": now.Unix(),
		"exp": now.Add(time.Hour * 1).Unix(),
	})

	secretKey := []byte(config.App.Server.Secret)
	return token.SignedString(secretKey)

}

func Decode() {}
