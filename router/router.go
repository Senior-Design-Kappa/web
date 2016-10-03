package router

import (
	"fmt"
  // "html/template"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/Senior-Design-Kappa/web/auth"
	"github.com/Senior-Design-Kappa/web/config"
	"github.com/Senior-Design-Kappa/web/logic"
)

type Server struct {
	*http.Server
	logic  logic.Logic
	Config config.Config
}

func NewServer(conf config.Config, logic logic.Logic, auth auth.Auth) *Server {
	r := mux.NewRouter()
  gets := r.Methods("GET").Subrouter()

  gets.HandleFunc("/", HomeHandler)
	gets.HandleFunc("/health", auth.DoAuth(health))

  auth.AddMountPath(r)
  stack := auth.CreateRouter(r)

	s := &Server{
		Server: &http.Server{
			Handler:      stack,
			Addr:         conf.Addr,
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		},
		logic:  logic,
		Config: conf,
	}

	return s
}

func health(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "OK")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Home")
  // t, _ := template.ParseFiles("templates/login.html.tpl")
  // t.Execute(w, nil)
}
