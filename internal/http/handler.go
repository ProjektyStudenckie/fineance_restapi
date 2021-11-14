package http

import (
	auth2 "ApiRest/internal/auth"
	mongo2 "ApiRest/internal/mongo"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
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
	h.Router.HandleFunc("/user", mongo2.CreateUserEndpoint).Methods("POST")
	h.Router.HandleFunc("/user/{id}", mongo2.GetUserEndpoint).Methods("GET")
	h.Router.HandleFunc("/login/{password}/{username}", auth2.Login).Methods("POST")
	h.Router.HandleFunc("/test/{test}", auth2.Test).Methods("GET")
}

