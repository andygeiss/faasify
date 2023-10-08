package main

import (
	"flag"
	"log"

	"github.com/andygeiss/faasify/internal/account"
	"github.com/andygeiss/faasify/internal/config"
	"github.com/andygeiss/faasify/internal/http/server"
)

func main() {
	// flags
	appName := flag.String("app.name", "faasify", "your app name")
	domain := flag.String("domain", "localhost", "your.domain")
	mode := flag.String("mode", "development", "development|production")
	url := flag.String("url", "http://localhost:3000", "remote server url")
	flag.Parse()
	// init config
	log.Printf("app.name: %s", *appName)
	log.Printf("domain:   %s", *domain)
	log.Printf("mode:     %s", *mode)
	log.Printf("url:      %s", *url)
	accountAccess := account.NewFileAccess("data/accounts.json")
	cfg := &config.Config{
		AccountAccess: accountAccess,
		AppName:       *appName,
		Domain:        *domain,
		Mode:          *mode,
		Url:           *url,
	}
	// server
	srv := server.NewManager(cfg)
	srv.ListenAndServe()
	if err := srv.Error(); err != nil {
		log.Fatal(err)
	}
}
