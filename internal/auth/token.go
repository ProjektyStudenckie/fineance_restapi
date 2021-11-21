package auth

import (
	"ApiRest/internal/mongo"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

var SecretLogin = []byte("secretLogin")
var SecretAccess = []byte("secretAccess")
var SecretRefresh = []byte("secretRefresh")

func GenerateLoginToken(user mongo.User) (map[string]string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = 1
	claims["name"] = user.Password

	t, err := token.SignedString(SecretLogin)
	if err != nil {
		return nil, err
	}
	return map[string]string{"login_token": t}, nil
}

func GenerateAccessToken(user mongo.User) (map[string]string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = 1
	claims["name"] = user.Username
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

	t, err := token.SignedString(SecretAccess)
	if err != nil {
		return nil, err
	}
	return map[string]string{"access_token": t}, nil
}


func GenerateRefreshToken(user mongo.User) (map[string]string, error) {
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["name"] = user.Username
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	rt, err := refreshToken.SignedString(SecretRefresh)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"refresh_token": rt,
	}, nil

}

func MethodAuth(original func(w http.ResponseWriter, r *http.Request), secret []byte) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {

			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return secret, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				original(w, r)
			}
		} else {

			fmt.Fprintf(w, "Not Authorized")
		}
	}
}
