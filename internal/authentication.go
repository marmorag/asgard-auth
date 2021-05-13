package internal

import (
	"fmt"
	"net/http"
)

type BasicAuthFunc func(username, password string) bool

func (f BasicAuthFunc) RequireAuth(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", `Basic realm="Authorization Required"`)
	w.WriteHeader(401)
}

func (f BasicAuthFunc) Authenticate(r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	return ok && f(username, password)
}

func HandleAuthenticate(w http.ResponseWriter, r *http.Request) {
	username, password := "foo", "bar"
	var authenticatedUser string

	f := BasicAuthFunc(func(user, pass string) bool {
		authenticatedUser = user
		return username == user && password == pass
	})

	if !f.Authenticate(r) {
		f.RequireAuth(w)
		writeLog(fmt.Sprintf("user '%s' failed to authenticate", authenticatedUser))
	} else {
		w.WriteHeader(200)
		writeLog(fmt.Sprintf("user '%s' authenticated", authenticatedUser))
	}
}