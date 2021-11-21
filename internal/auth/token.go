package auth

import (
	"ApiRest/internal/mongo"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var secrret = []byte("secret")

func GenerateToken(user mongo.User) (map[string]string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = 1
	claims["name"] = user.Password

	t, err := token.SignedString(secrret)
	if err != nil {
		return nil, err
	}
	return map[string]string{"access_token": t}, nil
}

func GenerateRefreshToken(user mongo.User) (map[string]string, error) {
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["name"] = user.Password
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	rt, err := refreshToken.SignedString(secrret)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"refresh_token": rt,
	}, nil

}
