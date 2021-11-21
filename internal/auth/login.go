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
	collection := mongo2.Client.Database("TestDB").Collection("Users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	_ = collection.FindOne(ctx, mongo2.User{Username: username}).Decode(&user)
	if pass == user.Password {
		tokens, err := GenerateTokenPair(user)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
			return
		}

		json.NewEncoder(response).Encode(tokens)
	}
}


func Register(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var user mongo2.User
	_ = json.NewDecoder(request.Body).Decode(&user)
	collection := mongo2.Client.Database("TestDB").Collection("Users")
	ctx,_ := context.WithTimeout(context.Background(), 5*time.Second)


	cur := collection.FindOne(ctx, mongo2.User{Username: user.Username})

	json.NewEncoder(response).Encode(cur)
	result, _ := collection.InsertOne(ctx, user)

	json.NewEncoder(response).Encode(result)
}

