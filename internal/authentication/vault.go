package authentication

import (
	"fmt"
	"github.com/tg123/go-htpasswd"
	"marmorag/asgard-auth/internal"
)

type VaultAuthenticationHandler struct {
	config internal.Configuration
	users map[string]User
}

func fetchUsers() map[string]User {
	userMap := make(map[string]User)
	users := internal.GetUsers()

	for user, _ := range users {
		userMap[user] = User{
			Username: user,
			Password: fmt.Sprintf("%s", users[user]),
		}
	}

	return userMap
}

func NewVaultAuthenticationHandler(configuration internal.Configuration) VaultAuthenticationHandler {
	handler := VaultAuthenticationHandler{config: configuration}
	handler.users = fetchUsers()

	return handler
}

func (e VaultAuthenticationHandler) AuthenticateUser(username string, password string) bool {
	if len(e.users) == 0 {
		err := "no vault users found"

		internal.WriteLog(err)
		panic(err)
	}

	user, keyExist := e.users[username]

	if !keyExist {
		return false
	}

	encoded, err := htpasswd.AcceptMd5(user.Password)

	if err != nil {
		internal.WriteLog(err.Error())
		panic(err.Error())
	}

	return encoded.MatchesPassword(password)
}