package auth

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/justinas/nosurf"

	"gopkg.in/authboss.v0"
	_ "gopkg.in/authboss.v0/auth"
	_ "gopkg.in/authboss.v0/register"
)

type Auth struct {
	ab *authboss.Authboss
}

func NewAuth() (Auth, error) {
	ab := setupAuthboss()
	a := Auth{
		ab: ab,
	}
	return a, nil
}

func (a Auth) DoAuth(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if u, err := a.ab.CurrentUser(w, r); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else if u == nil {
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			f.ServeHTTP(w, r)
		}
	}
}

func (a Auth) AddMountPath(r *mux.Router) {
	r.PathPrefix(a.ab.MountPath).Handler(a.ab.NewRouter())
}

func (a Auth) CreateRouter(r *mux.Router) http.Handler {
	return alice.New(nosurfing, a.ab.ExpireMiddleware).Then(r)
}

func setupAuthboss() *authboss.Authboss {
	ab := authboss.New()
	ab.Storer = NewDbUserStorer()
	ab.CookieStoreMaker = NewCookieStorer
	ab.SessionStoreMaker = NewSessionStorer
	ab.XSRFName = xsrfName
	ab.XSRFMaker = func(_ http.ResponseWriter, r *http.Request) string {
		return nosurf.Token(r)
	}
	ab.MountPath = "/auth"
	ab.ViewsPath = "./templates"
	if err := ab.Init(); err != nil {
		log.Fatal(err)
	}
	return ab
}

func nosurfing(h http.Handler) http.Handler {
	surfing := nosurf.New(h)
	surfing.SetFailureHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Failed to validate XSRF Token:", nosurf.Reason(r))
		w.WriteHeader(http.StatusBadRequest)
	}))
	return surfing
}
