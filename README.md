# Dalcs API
This repo contains the backend code for [dalcs-www](github.com/dalcs/dalcs-www).

## Getting started
To run the rest api server, run
```
go build -o api main.go
./api rest
```

## Routes

| Route        | Params           | Description  |
| ------------- |:-------------:| -----:|
| POST `/v1/invite/email` | to | Sends an email with a verification code to the `to` param. This route rate limits based on ip to 10 emails per hour. |
| POST `/v1/invite/verify` | code, email | Verifies the param `code` with the specified `email` param. It returns a JWT token with code & email encoded. This route rate limits 5req/min. |
| POST `/v1/invite/invite` | token, github_user | This route will verify the token and call github's api to invite the specified `github_user`. |

## Structure
- [main.go](main.go): Parse args and run appropriate cli command
- [internal/cmd](internal/cmd/rest): CLI commands for the rest service
- [internal/handlers](internal/handlers): Holds the rest api handlers
- [internal/rest](internal/rest): The actual rest api logic