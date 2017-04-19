package router

import (
	"fmt"
	"net/http"
  "net/url"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"github.com/Senior-Design-Kappa/web/auth"
	"github.com/Senior-Design-Kappa/web/config"
	"github.com/Senior-Design-Kappa/web/logic"
)

type Server struct {
	*http.Server
  Config *ServerConfig
}

type ServerConfig struct {
	logic  logic.Logic
	config config.Config
  auth   auth.Auth
}

// TODO: this seems kinda janky..?
// var serverConf *config.Config
// var serverAuth *auth.Auth
// var serverLogic *logic.Logic

func NewServer(conf config.Config, logic logic.Logic, auth auth.Auth) *Server {
	r := mux.NewRouter()
  sc := &ServerConfig {
    logic: logic,
    config: conf,
    auth: auth,
  }
	// serverConf = &conf
	// serverAuth = &auth

	// GET request handlers
	gets := r.Methods("GET").Subrouter()
	gets.HandleFunc("/", sc.HomeHandler)
	gets.HandleFunc("/health", auth.DoAuth(health))
	gets.HandleFunc("/room/{id}", sc.RoomHandler)

  // API endpoints
  gets.HandleFunc("/api/createRoom", sc.CreateRoomHandler);

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
    Config: sc,
	}

	return s
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func (s ServerConfig) HomeHandler(w http.ResponseWriter, r *http.Request) {
	gets := r.URL.Query()
	var showLogin = ""
	if len(gets["showLogin"]) > 0 {
		showLogin = "true"
	}

  data := Data{
    Title: "Kappa",
    ShowLogin: showLogin,
  }
	s.RenderHeaderFooterTemplate(w, r, data, "templates/index.html")
}

func (s ServerConfig) RoomHandler(w http.ResponseWriter, r *http.Request) {
  // Parse as int64
  roomIdStr := mux.Vars(r)["id"]
	roomId, err := strconv.ParseInt(roomIdStr, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Invalid room ID! (%s)", err)
		return
	}
  videoId, err := s.logic.GetVideoId(roomId)

  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "Invalid video link! (%s)", err)
    return
  }

  data := Data{
    WebsocketAddr: "ws://" + s.config.SyncAddr + "/connect/",
    RoomId:        roomIdStr,
    Title:         "Room " + roomIdStr,
    VideoId:     videoId,
  }
	s.RenderHeaderFooterTemplate(w, r, data, "templates/room.html")
}

func (s ServerConfig) CreateRoomHandler(w http.ResponseWriter, r *http.Request) {
  // Get user data
  user, err := s.auth.GetCurrentUser(w, r)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "Authentication error! (%s)", err)
    return
  }
  userId, err := s.auth.GetIdFromUser(user)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "Authentication error! (%s)", err)
    return
  }

  // Get video link data
	gets := r.URL.Query()
	if len(gets["videoLink"]) == 0 {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "No video link!")
    return
	}

  videoUrl, err := url.Parse(gets["videoLink"][0])
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "Video link parse error! (%s)", err)
    return
  }

  videoId := videoUrl.Query()["v"][0]

  // Try to create room
  roomId, err := s.logic.CreateRoom(userId, videoId)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "Room creation error! (%s)", err)
    return
  }
  fmt.Fprintf(w, "%d", roomId)
}
