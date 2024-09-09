package jwt

import (
	"fmt"
	"time"

	jwtcore "github.com/golang-jwt/jwt/v5"
	"github.com/rootspyro/50BEERS/config"
)


func Encode(email string, hours time.Duration) (string, error) {

	now := time.Now()

	token := jwtcore.NewWithClaims(jwtcore.SigningMethodHS256, jwtcore.MapClaims{
		"sub": email, // subject
		"iat": now.Unix(),
		"exp": now.Add(time.Hour * hours).Unix(),
	})

	secretKey := []byte(config.App.Server.Secret)
	return token.SignedString(secretKey)

}

func SignToken(email string) (string, error) {
	return Encode(email, 1)
}

func SignRefreshToken(email string) (string, error) {
	return Encode(email, 12)
}

func Decode(tokenStr string) error {

	secretKey := []byte(config.App.Server.Secret)

	token, err := jwtcore.Parse(tokenStr, func(t *jwtcore.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwtcore.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	fmt.Println(token)

	return err
}
