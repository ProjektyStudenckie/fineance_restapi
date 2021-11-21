package auth

import (
	"ApiRest/internal/mongo"
	"time"
	"github.com/dgrijalva/jwt-go"
)

var secrret = []byte("secret")

func GenerateTokenPair(user mongo.User) (map[string]string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = 1
	claims["name"] = user.Username
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	t, err := token.SignedString(secrret)
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	rt, err := refreshToken.SignedString(secrret)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  t,
		"refresh_token": rt,
	}, nil
}