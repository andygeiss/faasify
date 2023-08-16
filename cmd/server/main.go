package main

import (
	"flag"
	"log"

	"github.com/andygeiss/faasify/internal/account"
	"github.com/andygeiss/faasify/internal/http/server"
)

func main() {
	appName := flag.String("app.name", "faasify", "your app name")
	domain := flag.String("domain", "localhost", "your.domain")
	mode := flag.String("mode", "development", "development|production")
	url := flag.String("url", "https://localhost:3000", "remote server url")
	flag.Parse()
	accountAccess := account.NewFileAccess("data/accounts.json")
	srv := server.NewManager().
		WithAccountAccess(accountAccess).
		WithAppName(*appName).
		WithDomain(*domain).
		WithMode(*mode).
		WithUrl(*url)
	log.Printf("app.name: %s", *appName)
	log.Printf("domain:   %s", *domain)
	log.Printf("mode:     %s", *mode)
	log.Printf("url:      %s", *url)
	srv.ListenAndServe()
	if err := srv.Error(); err != nil {
		log.Fatal(err)
	}
}
