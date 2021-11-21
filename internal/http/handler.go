package http

import (
	auth2 "ApiRest/internal/auth"
	webSockets2 "ApiRest/internal/websockets"
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
	h.Router.HandleFunc("/register", auth2.Register).Methods("GET")
	h.Router.HandleFunc("/login", auth2.Login).Methods("GET")
	h.Router.HandleFunc("/test/{test}", auth2.Test).Methods("GET")
	h.Router.HandleFunc("/stats", webSockets2.TestSocket)
	http.ListenAndServe(":1332", h.Router)
}


