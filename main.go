package main

import (
	"os"

	"github.com/dalcs/dalcs-api/internal/cmd/rest"
	"github.com/jessevdk/go-flags"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	p := flags.NewParser(nil, flags.Default)
	rest.Register(p)

	if _, err := p.Parse(); err != nil {
		os.Exit(1)
	}
}
