package email

import (
	"fmt"
	"os"

	"github.com/matcornic/hermes/v2"
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Email struct {
	to   string
	code string
}

// New creates a new email struct
func New(to, code string) *Email {
	return &Email{
		to:   to,
		code: code,
	}
}

// Send sends the email to specified recipient
func (e *Email) Send() (*rest.Response, error) {
	if !e.Validate() {
		return nil, fmt.Errorf("Unspecified required email fields")
	}

	// Create the email template
	config := hermes.Hermes{
		Product: hermes.Product{
			Name:        "@dalcs",
			Link:        "https://dalcs.mackboudreau.com/",
			Copyright:   "Unofficial Dalhousie CS GitHub Organization",
			Logo:        "https://svgur.com/i/GwA.svg",
			TroubleText: "If you’re having trouble with the invite code, send an email to hi@mackboudreau.com",
		},
	}
	email := hermes.Email{
		Body: hermes.Body{
			Intros: []string{
				"You’re one step closer to being invited to github.com/dalcs! You’ll have to use the code below to verify you own this email address.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "Please copy your verification code:",
					InviteCode:   e.code,
				},
			},
			Outros: []string{
				"Didn’t request this email? That’s fine, just ignore it.",
			},
		},
	}

	emailBody, err := config.GenerateHTML(email)
	if err != nil {
		return nil, err
	}
	emailText, err := config.GeneratePlainText(email)
	if err != nil {
		return nil, err
	}

	// Send the email via sendgrid
	from := mail.NewEmail("Dalcs GitHub Organization", "no-reply@mackboudreau.com")
	to := mail.NewEmail("Mack boudreau", e.to)
	subject := "Verify your Dalhousie email address"
	message := mail.NewSingleEmail(from, subject, to, emailText, emailBody)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Validate checks whether the fields have been properly initalized
func (e *Email) Validate() bool {
	if e.to == "" {
		return false
	} else if e.code == "" {
		return false
	}
	return true
}
