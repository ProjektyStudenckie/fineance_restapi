package http

import (
	"ApiRest/internal/auth"
	"ApiRest/internal/mongo"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type Handler struct{
	Router *mux.Router
}

func NewHandler() *Handler{
	return &Handler{}
}

func (h *Handler) SetupRoutes(){
	fmt.Println("Setting Up Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health",func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"alive")
	})
}

func Login(response http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	username := params["username"]
	password := params["password"]
	var user mongo.User
	collection := mongo.Client.Database("TestDB").Collection("Users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	_ = collection.FindOne(ctx, mongo.User{Username: username}).Decode(&user)
	if  password == user.Password {
		tokens, err := auth.GenerateTokenPair()
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
			return
		}

		json.NewEncoder(response).Encode(tokens)
	}
}

func Test(response http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	test := params["test"]
		json.NewEncoder(response).Encode(test)
}
