package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"marmorag/asgard-auth/internal/authentication"
	"net/http"
	"os"
)

func executeRootCommand() {
	fmt.Println("Authentication is running...")

	http.Handle("/", http.HandlerFunc(authentication.HandleAuthenticate))
	http.ListenAndServe(":3000", nil)
}

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "Basic authentication server to use with Vault",
	Run: func(cmd *cobra.Command, args []string) {
		executeRootCommand()
	},
}

func init() {
	// init things for root command
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
