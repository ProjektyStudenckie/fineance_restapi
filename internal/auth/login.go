package auth

import (
	mongo2 "ApiRest/internal/mongo"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

func Login(response http.ResponseWriter, request *http.Request) {
	username, pass, _ := request.BasicAuth()
	var user mongo2.User
	collection := mongo2.DataBaseCon.Client.Database("TestDB").Collection("Users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	cur := collection.FindOne(ctx, mongo2.User{Username: username}).Decode(&user)
	if cur == nil {
		token, err := GenerateToken(mongo2.User{Username: username, Password: pass})
		if token["access_token"] == user.Password {
			if err != nil {
				response.WriteHeader(http.StatusInternalServerError)
				response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
				return
			}
			tokenRT, _ := GenerateRefreshToken(mongo2.User{Username: username, Password: pass})
			json.NewEncoder(response).Encode(tokenRT)
		} else {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{ "message": "wrong password" }`))
			return
		}
	} else {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "no user with this id" }`))
		return
	}
}

func Register(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var user mongo2.User
	_ = json.NewDecoder(request.Body).Decode(&user)
	collection := mongo2.DataBaseCon.Client.Database("TestDB").Collection("Users")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	cur := collection.FindOne(ctx, mongo2.User{Username: user.Username})

	if cur.Err() != nil {
		token, _ := GenerateToken(user)
		user.Password = token["access_token"]
		_, err := collection.InsertOne(ctx, user)
		if err == nil {
			json.NewEncoder(response).Encode(`{ "message": "Username Registered" }`)
		}
	} else {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "Username Taken" }`))
		return
	}

}
