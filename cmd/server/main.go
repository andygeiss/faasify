package main

import (
	"flag"
	"log"

	"github.com/andygeiss/faasify/internal/http/server"
)

func main() {
	domain := flag.String("domain", "faasify.dev", "your.domain")
	mode := flag.String("mode", "", "prod for production")
	url := flag.String("url", "https://faasify.dev", "remote server url")
	flag.Parse()
	log.Printf("domain: %s", *domain)
	log.Printf("mode:   %s", *mode)
	log.Printf("url:    %s", *url)
	if err := server.ListenAndServe(*domain, *mode, *url); err != nil {
		log.Fatal(err)
	}
}
