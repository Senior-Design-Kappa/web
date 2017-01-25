package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (a Auth) CreateAuthSubRouter(r *mux.Router) {
	r.HandleFunc("/register", a.registerHandler)
	r.HandleFunc("/login", a.loginHandler)
}

func responseError(w http.ResponseWriter) {
	fmt.Fprintf(w, "{error:true}")
}

func (a Auth) registerHandler(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	if len(vars["username"]) != 1 || len(vars["password"]) != 1 || len(vars["email"]) != 1 {
		log.Printf("Bad registration!\n")
		responseError(w)
		return
	}
	username := vars["username"][0]
	password := vars["password"][0]
	email := vars["email"][0]
	if username == "" || password == "" || email == "" {
		log.Printf("Bad registration!\n")
		responseError(w)
		return
	}
	err := a.RegisterNewUser(username, password, email)
	if err != nil {
		log.Printf("%+v\n", err)
		responseError(w)
		return
	}
	fmt.Fprintf(w, "{ok:true}")
}

func (a Auth) loginHandler(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	if len(vars["username"]) != 1 || len(vars["password"]) != 1 {
		log.Printf("Bad login!\n")
		responseError(w)
		return
	}
	username := vars["username"][0]
	password := vars["password"][0]
	token, err := a.LoginUser(username, password)
	if err != nil {
		log.Printf("%+v\n", err)
		responseError(w)
		return
	}
	data := make(map[string]string)
	data["token"] = token
	reply, err := json.Marshal(data)
	if err != nil {
		log.Printf("%+v\n", err)
		responseError(w)
		return
	}
	fmt.Fprintf(w, "%s", reply)
}
