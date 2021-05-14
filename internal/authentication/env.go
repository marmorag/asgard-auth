package authentication

import (
	"encoding/base64"
	"github.com/tg123/go-htpasswd"
	"marmorag/asgard-auth/internal"
	"strings"
)

type EnvironmentAuthenticationHandler struct {
	config internal.Configuration
	users map[string]User
}

func NewEnvironmentAuthenticationHandler(configuration internal.Configuration) EnvironmentAuthenticationHandler {
	handler := EnvironmentAuthenticationHandler{config: configuration}
	handler.users = parseAuthString(configuration.EnvironmentAuthString)

	return handler
}

func parseAuthString(s string) map[string]User {
	users := make(map[string]User)

	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		internal.WriteLog(err.Error())
		panic(err.Error())
	}

	// Auth String Format
	// user:password\nuser:password ...
	strData := string(data)
	strUsers := strings.Split(strData, "\n")

	for _, strUser := range strUsers {
		userCredential := strings.Split(strUser, ":")

		users[userCredential[0]] = User{
			Username: userCredential[0],
			Password: userCredential[1],
		}
	}

	return users
}

func (e EnvironmentAuthenticationHandler) AuthenticateUser(username string, password string) bool {
	if e.config.EnvironmentAuthString == "" {
		err := "no environment authentication string found"

		internal.WriteLog(err)
		panic(err)
	}

	user, userExist := e.users[username]

	if !userExist {
		return false
	}

	encoded, err := htpasswd.AcceptMd5(user.Password)

	if err != nil {
		internal.WriteLog(err.Error())
		panic(err.Error())
	}

	return encoded.MatchesPassword(password)
}
