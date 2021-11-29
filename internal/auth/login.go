package auth

import (
	mongo2 "ApiRest/internal/mongo"
	"context"
	"encoding/json"
	"fmt"
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
		token, err := GenerateLoginToken(mongo2.User{Username: username, Password: pass})
		if token["login_token"] == user.Password {
			if err != nil {
				response.WriteHeader(http.StatusInternalServerError)
				response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
				return
			}
			tokenAccess, _ := GenerateAccessToken(mongo2.User{Username: username, Password: pass})
			tokenRT, _ := GenerateRefreshToken(mongo2.User{Username: username, Password: pass})
			user.RT = tokenRT["refresh_token"]
			cur2,cur3 :=collection.ReplaceOne(ctx,mongo2.User{Username: username},
				user,
			)
			fmt.Println(cur2)
			fmt.Println(cur3)
			json.NewEncoder(response).Encode(MergeJSONMaps(tokenAccess, tokenRT))
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
		token, _ := GenerateLoginToken(user)
		user.Password = token["login_token"]
		tokenAccess, _ := GenerateAccessToken(user)
		rt,_:=GenerateRefreshToken(user)
		user.RT = rt["refresh_token"]
		_, err := collection.InsertOne(ctx, user)
		if err == nil {
			json.NewEncoder(response).Encode(MergeJSONMaps(tokenAccess,rt))
		}
	} else {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "Username Taken" }`))
		return
	}
}

func Refresh(response http.ResponseWriter, request *http.Request) {
	if request.Header["Token"] != nil {
		var user mongo2.User
		collection := mongo2.DataBaseCon.Client.Database("TestDB").Collection("Users")
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

		cur := collection.FindOne(ctx, mongo2.User{RT: request.Header["Token"][0]}).Decode(&user)
		if cur == nil {
			tokenAccess, _ := GenerateAccessToken(user)
			tokenRT, _ := GenerateRefreshToken(user)
			user.RT = tokenRT["refresh_token"]
			collection.ReplaceOne(ctx,mongo2.User{Username: user.Username},
				user,
			)
			json.NewEncoder(response).Encode(MergeJSONMaps(tokenAccess, tokenRT))
		}
	}
}

func MergeJSONMaps(maps ...map[string]string) (result map[string]string) {
	result = make(map[string]string)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}
