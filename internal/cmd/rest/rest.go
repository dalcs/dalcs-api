package rest

import (
	"log"
	"net/http"

	api "github.com/dalcs/dalcs-api/internal/handlers"
	"github.com/jessevdk/go-flags"
)

const port = "8080"

// Register the rest commands
func Register(p *flags.Parser) {
	_, err := p.AddCommand("rest", "start the rest api service", "", &command{})
	if err != nil {
		log.Fatalln(err)
	}
}

type command struct{}

func (command) Execute(args []string) error {
	log.Printf("Server started on %s\n", port)
	http.HandleFunc("/v1/invite/email", api.InviteEmailHandler)
	http.HandleFunc("/v1/invite/verify", api.InviteVerifyHandler)
	http.HandleFunc("/v1/invite/invite", api.InviteInviteHandler)
	return http.ListenAndServe(":"+port, nil)
}
