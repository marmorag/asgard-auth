package internal

import (
	"fmt"
	vault "github.com/hashicorp/vault/api"
	"os"
)

var vaultClient *vault.Client

func init() {
	appConfig := GetConfig()
	if appConfig.AuthenticationMode != VaultAuthMode {
		return
	}

	vaultConfig := vault.DefaultConfig()
	vaultConfig.Address = os.Getenv("VAULT_ADDR")

	var err error
	vaultClient, err = vault.NewClient(vaultConfig)

	if err != nil {
		WriteLog(err.Error())
		return
	}

	vaultClient.SetToken(os.Getenv("VAULT_TOKEN"))
}

func GetClient() *vault.Client {
	return vaultClient
}

func GetUsers() map[string]interface{} {
	secret, err := vaultClient.Logical().Read("secret/data/asgard-auth/users")

	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}

	usersSecret, ok := secret.Data["data"].(map[string]interface{})

	if !ok {
		panic("no data found")
	}

	return usersSecret
}

func PingVault() bool {
	secret, err := vaultClient.Logical().Read("secret/data/asgard-auth")

	if err != nil {
		fmt.Println(err)
		return false
	}

	m, ok := secret.Data["data"].(map[string]interface{})

	if !ok {
		fmt.Printf("%T %#v\n", secret.Data["data"], secret.Data["data"])
		return false
	}

	return m["ping"] == "ok"
}
