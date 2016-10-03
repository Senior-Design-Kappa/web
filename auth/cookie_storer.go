// Based on authboss-sample

package auth

import (
  "fmt"
	"net/http"
	"time"

  "github.com/gorilla/securecookie"

  "gopkg.in/authboss.v0"
)

type CookieStorer struct {
	w http.ResponseWriter
	r *http.Request
  cookieStore *securecookie.SecureCookie
}

func NewCookieStorer(w http.ResponseWriter, r *http.Request) authboss.ClientStorer {
  return &CookieStorer {w, r, securecookie.New(cookieStoreKey, nil)}
}

func (s CookieStorer) Get(key string) (string, bool) {
  cookie, err := s.r.Cookie(key)
	if err != nil {
		return "", false
	}

	var value string
	err = s.cookieStore.Decode(key, cookie.Value, &value)
	if err != nil {
		return "", false
	}

	return value, true
}

func (s CookieStorer) Put(key, value string) {
	encoded, err := s.cookieStore.Encode(key, value)
	if err != nil {
		fmt.Println(err)
	}

	cookie := &http.Cookie{
		Expires: time.Now().UTC().AddDate(1, 0, 0),
		Name:    key,
		Value:   encoded,
		Path:    "/",
	}
	http.SetCookie(s.w, cookie)
}

func (s CookieStorer) Del(key string) {
	cookie := &http.Cookie{
		MaxAge: -1,
		Name:   key,
		Path:   "/",
	}
	http.SetCookie(s.w, cookie)
}
