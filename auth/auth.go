package auth

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/justinas/nosurf"

	"github.com/Senior-Design-Kappa/web/config"
)

type Auth struct {
	db          *sql.DB
	MountPath   string
	TokenSecret []byte
}

func NewAuth(conf config.Config) (*Auth, error) {
	db, err := sql.Open("sqlite3", "./db/users.db")
	if err != nil {
		log.Printf("error: could not open db connection (%+v)\n", err)
		return nil, err
	}
	// TODO: change token
	a := &Auth{
		db:          db,
		MountPath:   "/auth/",
		TokenSecret: []byte("testingtestingtestingtestingtestingtestingtestingtestingtestingtesting"),
	}
	return a, nil
}

func (a Auth) AddMountPath(r *mux.Router) {
	sr := r.PathPrefix(a.MountPath).Subrouter()
	a.CreateAuthSubRouter(sr)
}

func (a Auth) WrapXSRFRouter(r *mux.Router) http.Handler {
	return alice.New(nosurfing).Then(r)
}

func nosurfing(h http.Handler) http.Handler {
	surfing := nosurf.New(h)
	surfing.SetFailureHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Failed to validate XSRF Token:", nosurf.Reason(r))
		w.WriteHeader(http.StatusBadRequest)
	}))
	return surfing
}

func (a Auth) DoAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := a.GetCurrentUser(w, r)
		if err != nil || user == "" {
			log.Printf("%+v\n", err)
			return
		}
		h(w, r)
	}
}

func (a Auth) GetCurrentUser(w http.ResponseWriter, r *http.Request) (string, error) {
	cookie, err := r.Cookie("JWT_TOKEN")
	if err != nil {
		return "", err
	}
	user, err := a.GetUserFromToken(cookie.Value)
	return user, err
}
