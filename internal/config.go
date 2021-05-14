package internal

import "os"

type AuthMode string

const (
	EnvironmentAuthMode = "env"
	VaultAuthMode       = "vault"
)

type Configuration struct {
	AuthenticationMode AuthMode

	EnvironmentAuthString string

	VaultAddress string
	VaultToken string
}

var config Configuration

func GetConfig() Configuration {
	return config
}

func init() {
	config = Configuration{
		AuthenticationMode: AuthMode(os.Getenv("AUTH_MODE")),
	}

	if config.AuthenticationMode == EnvironmentAuthMode {
		config.EnvironmentAuthString = os.Getenv("AUTH_STRING")
	}

	if config.AuthenticationMode == VaultAuthMode {
		config.VaultAddress = os.Getenv("VAULT_ADDR")
		config.VaultToken = os.Getenv("VAULT_TOKEN")
	}
}
