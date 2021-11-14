package auth

import (
	mongo2 "ApiRest/internal/mongo"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func Login(response http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	username := params["username"]
	password := params["password"]
	var user mongo2.User
	collection := mongo2.Client.Database("TestDB").Collection("Users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	_ = collection.FindOne(ctx, mongo2.User{Username: username}).Decode(&user)
	if  password == user.Password {
		tokens, err := GenerateTokenPair(user)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
			return
		}

		json.NewEncoder(response).Encode(tokens)
	}
}