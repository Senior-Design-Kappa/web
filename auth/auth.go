package auth

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/justinas/nosurf"

	"github.com/Senior-Design-Kappa/web/config"
)

type Auth struct {
	db        *sql.DB
	MountPath string
}

func NewAuth(conf config.Config) (Auth, error) {
	db, err := sql.Open("sqlite3", "./db/users.db")
	if err != nil {
		log.Printf("error: could not open db connection (%+v)\n", err)
		return nil, err
	}
	a := Auth{
		db:        db,
		MountPath: "/auth/",
	}
	return a, nil
}

func (a Auth) AddMountPath(r *mux.Router) {
	r.PathPrefix(a.MountPath).Handler(a.CreateAuthSubRouter())
}

func (a Auth) WrapXSRFRouter(r *mux.Router) http.Handler {
	return alice.New(nosurfing, a.ab.ExpireMiddleware).Then(r)
}

func nosurfing(h http.Handler) http.Handler {
	surfing := nosurf.New(h)
	surfing.SetFailureHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Failed to validate XSRF Token:", nosurf.Reason(r))
		w.WriteHeader(http.StatusBadRequest)
	}))
	return surfing
}
