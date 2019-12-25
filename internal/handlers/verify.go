package api

import (
	"net/http"
)

func InviteVerifyHandler(res http.ResponseWriter, req *http.Request) {
	// Do some rate limiting here...

	// Verify code + email

	// Generate signed JWT

	// Return token

	return
}
