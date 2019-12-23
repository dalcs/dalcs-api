package rest

import (
	"encoding/json"
	"net/http"

	email "github.com/dalcs/dalcs-api/internal/rest"
)

func InviteEmailHandler(res http.ResponseWriter, req *http.Request) {
	e := email.New("mack.boudreau2@gmail.com", "99999")
	resp, err := e.Send()
	if err != nil {
		res.WriteHeader(500)
		res.Write([]byte("There has been a fatal error..."))
		return
	}
	respStr, err := json.Marshal(resp)
	if err != nil {
		res.WriteHeader(500)
		res.Write([]byte("There has been a fatal error..."))
		return
	}
	res.Header().Add("content-type", "application/json")
	res.Write([]byte(respStr))
	return
}
