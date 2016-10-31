package router

import (
	"fmt"
	"html/template"
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

func NewServer(conf config.Config, logic logic.Logic, auth auth.Auth) *Server {
	r := mux.NewRouter()
	serverConf = &conf

	// GET request handlers
	gets := r.Methods("GET").Subrouter()
	gets.HandleFunc("/", HomeHandler)
	gets.HandleFunc("/health", auth.DoAuth(health))
	gets.HandleFunc("/room/{id}", RoomHandler)

	// Static handlers
	gets.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir(conf.ClientPath+"css/"))))
	gets.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir(conf.ClientPath+"js/"))))

	// Auth stuff
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
	t, err := template.ParseFiles(
		serverConf.ClientPath+"templates/index.html",
		serverConf.ClientPath+"templates/header.html",
		serverConf.ClientPath+"templates/footer.html",
	)
	if err != nil {
		fmt.Fprintf(w, "Error with template! (%s)", err)
		return
	}
	data := struct {
		Title string
	}{
		"Kappa",
	}
	t.Execute(w, data)
}

func RoomHandler(w http.ResponseWriter, r *http.Request) {
	roomId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		fmt.Fprintf(w, "Invalid room ID!")
		return
	}
	t, err := template.ParseFiles(
		serverConf.ClientPath+"templates/room.html",
		serverConf.ClientPath+"templates/header.html",
		serverConf.ClientPath+"templates/footer.html",
	)
	if err != nil {
		fmt.Fprintf(w, "Error with template! (%s)", err)
		return
	}

	data := struct {
		WebsocketAddr string
		RoomId        int
		Title         string
	}{
		"ws://" + serverConf.SyncAddr + "/connect/",
		roomId,
		"Room " + strconv.Itoa(roomId),
	}
	t.Execute(w, data)
}
