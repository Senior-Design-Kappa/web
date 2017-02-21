package router

import (
	"fmt"
	"net/http"
	"strconv"
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


var serverConf *config.Config
var serverAuth *auth.Auth

func NewServer(conf config.Config, logic logic.Logic, auth auth.Auth) *Server {
	r := mux.NewRouter()
	serverConf = &conf
	serverAuth = &auth

	// GET request handlers
	gets := r.Methods("GET").Subrouter()
	gets.HandleFunc("/", HomeHandler)
	gets.HandleFunc("/health", auth.DoAuth(health))
	gets.HandleFunc("/room/{id}", RoomHandler)

	// Static handlers
	gets.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir(conf.ClientPath+"css/"))))
	gets.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir(conf.ClientPath+"js/"))))
	gets.PathPrefix("/imgs/").Handler(http.StripPrefix("/imgs/", http.FileServer(http.Dir(conf.ClientPath+"imgs/"))))

	// Auth stuff
	auth.AddMountPath(r)
	stack := auth.WrapXSRFRouter(r)

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
	gets := r.URL.Query()
	var showLogin = ""
	if len(gets["showLogin"]) > 0 {
		showLogin = "true"
	}

  data := Data{
    Title: "Kappa",
    ShowLogin: showLogin,
  }
	RenderHeaderFooterTemplate(w, r, data, "templates/index.html")
}

func RoomHandler(w http.ResponseWriter, r *http.Request) {
	roomId := mux.Vars(r)["id"]
	_, err := strconv.Atoi(roomId)
	if err != nil {
		fmt.Fprintf(w, "Invalid room ID! (%s)", err)
		return
	}

  data := Data{
  	WebsocketAddr: "ws://" + serverConf.SyncAddr + "/connect/",
  	RoomId:        roomId,
  	Title:         "Room " + roomId,
  }
	RenderHeaderFooterTemplate(w, r, data, "templates/room.html")
}
