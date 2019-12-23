package main

import (
	"net/http"

	"github.com/dalcs/dalcs-api/internal/cmd/rest"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// TODO: Move route handling into internal/cmd/rest & setup cmds
	http.HandleFunc("/v1/invite/email", rest.InviteEmailHandler)
	http.HandleFunc("/v1/invite/verify", rest.InviteVerifyHandler)
	http.HandleFunc("/v1/invite/invite", rest.InviteInviteHandler)
	http.ListenAndServe(":8080", nil)
}
