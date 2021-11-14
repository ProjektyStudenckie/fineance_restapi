package auth

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func Test(response http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	test := params["test"]
	json.NewEncoder(response).Encode(test)
}
