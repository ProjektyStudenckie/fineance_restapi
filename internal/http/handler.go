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
	h.Router.HandleFunc("/login", auth2.Login).Methods("Get")
	h.Router.HandleFunc("/test/{test}", auth2.MethodAuth(auth2.Test,auth2.SecretAccess)).Methods("Get")
	h.Router.HandleFunc("/refresh_access", auth2.MethodAuth(auth2.Refresh,auth2.SecretRefresh)).Methods("Get")
	h.Router.HandleFunc("/wallets", Wallet.GetSubWallets)
	h.Router.HandleFunc("/add_wallet", Wallet.AddWallet)
	h.Router.HandleFunc("/sub_wallets", Wallet.GetWallets)
	h.Router.HandleFunc("/stats", webSockets2.TestSocket)

	http.ListenAndServe(":1332", h.Router)
}


