package authentication

import (
	"fmt"
	"marmorag/asgard-auth/internal"
	"net/http"
)

type BasicAuthFunc func(username, password string) bool

type AuthenticationHandler interface {
	AuthenticateUser(username string, password string) bool
}

type User struct {
	Username string
	Password string
}

func buildAuthenticationHandler(configuration internal.Configuration) AuthenticationHandler {
	if configuration.AuthenticationMode == internal.EnvironmentAuthMode {
		return NewEnvironmentAuthenticationHandler(configuration)
	}

	if configuration.AuthenticationMode == internal.VaultAuthMode {
		return NewVaultAuthenticationHandler(configuration)
	}
	return nil
}

func (f BasicAuthFunc) RequireAuth(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", `Basic realm="Authorization Required"`)
	w.WriteHeader(401)
}

func (f BasicAuthFunc) Authenticate(r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	return ok && f(username, password)
}

func HandleAuthenticate(w http.ResponseWriter, r *http.Request) {
	config := internal.GetConfig()
	var authenticatedUser string

	authHandler := buildAuthenticationHandler(config)

	f := BasicAuthFunc(func(user, pass string) bool {
		authenticatedUser = user
		return authHandler.AuthenticateUser(user, pass)
	})

	if !f.Authenticate(r) {
		f.RequireAuth(w)
		internal.WriteLog(fmt.Sprintf("user '%s' failed to authenticate", authenticatedUser))
	} else {
		w.WriteHeader(200)
		internal.WriteLog(fmt.Sprintf("user '%s' authenticated", authenticatedUser))
	}
}
