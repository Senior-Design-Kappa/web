package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (a Auth) CreateAuthSubRouter() mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/register", a.registerHandler)
	return r
}

func responseError(w http.ResponseWriter) {
	fmt.Fprintf(w, "{error:true}")
}

func (a Auth) registerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	password := vars["password"]
	email := vars["email"]
	if username == "" || password == "" || email == "" {
		log.Printf("Bad registration!\n")
		responseError(w)
		return
	}
	err := a.RegisterNewUser(username, password, email)
	if err != nil {
		log.Printf("%s\n", err)
		return
	}
}
