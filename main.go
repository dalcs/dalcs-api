package main

import (
	"encoding/json"
	"net/http"

	"github.com/dalcs/dalcs-api/internal/cmd/rest"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("content-type", "application/json")
		user := struct {
			Username string
			Password string
		}{
			Username: "mack",
			Password: "secretpass",
		}

		userJSON, _ := json.Marshal(user)
		res.Write(userJSON)
	})
	http.HandleFunc("/email", rest.InviteEmailHandler)
	http.ListenAndServe(":8080", nil)
}
