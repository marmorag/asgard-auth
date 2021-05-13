package cmd

import (
	"fmt"
	"marmorag/asgard-auth/internal"
	"net/http"
)


func Execute() {
	fmt.Println("Authentication is running...")

	http.Handle("/", http.HandlerFunc(internal.HandleAuthenticate))
	http.ListenAndServe(":3000", nil)
}