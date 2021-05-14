package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"marmorag/asgard-auth/internal"
)

func vault() {
	var pingResponse string
	canPing := internal.PingVault()

	if canPing {
		pingResponse = "OK \xE2\x9C\x94"
	} else {
		pingResponse = "KO \xE2\x9C\x96"
	}

	fmt.Printf("ping to vault server : %s\n", pingResponse)
}

var vaultCmd = &cobra.Command{
	Use:   "vault",
	Short: "Test vault connection",
	Run: func(cmd *cobra.Command, args []string) {
		vault()
	},
}

func init() {
	rootCmd.AddCommand(vaultCmd)
}
