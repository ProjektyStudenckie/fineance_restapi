package http

import (
	"ApiRest/internal/Wallet"
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
	h.Router.HandleFunc("/register", auth2.Register).Methods("Post")
	h.Router.HandleFunc("/login", auth2.Login).Methods("GET")
	h.Router.HandleFunc("/test/{test}", auth2.MethodAuth(auth2.Test,auth2.SecretAccess)).Methods("GET")
	h.Router.HandleFunc("/refresh_access", auth2.MethodAuth(auth2.Refresh,auth2.SecretRefresh)).Methods("GET")
	h.Router.HandleFunc("/sub_wallets", auth2.MethodAuth(Wallet.GetSubWallets,auth2.SecretAccess)).Methods("GET")
	h.Router.HandleFunc("/add_wallet",auth2.MethodAuth( Wallet.AddWallet,auth2.SecretAccess)).Methods("POST")
	h.Router.HandleFunc("/wallets", auth2.MethodAuth(Wallet.GetWallets,auth2.SecretAccess)).Methods("GET")
	h.Router.HandleFunc("/add_sub_owner", auth2.MethodAuth(Wallet.AddSubOwner,auth2.SecretAccess)).Methods("POST")
	h.Router.HandleFunc("/updateWallet", auth2.MethodAuth(Wallet.UpdateWallet,auth2.SecretAccess)).Methods("POST")
	h.Router.HandleFunc("/remove_goal", auth2.MethodAuth(Wallet.RemoveGoal,auth2.SecretAccess)).Methods("POST")
	h.Router.HandleFunc("/add_goal", auth2.MethodAuth(Wallet.AddGoal,auth2.SecretAccess)).Methods("POST")
	h.Router.HandleFunc("/remove_sub_owner", auth2.MethodAuth(Wallet.RemoveSubOwner,auth2.SecretAccess)).Methods("POST")
	h.Router.HandleFunc("/add_remittance", auth2.MethodAuth(Wallet.AddRemittance,auth2.SecretAccess)).Methods("POST")
	h.Router.HandleFunc("/stats", auth2.MethodAuth(webSockets2.TestSocket,auth2.SecretAccess))

	http.ListenAndServe(":1332", h.Router)
}


