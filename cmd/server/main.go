package main

import (
	"log"

	"github.com/andygeiss/faasify/internal/http/server"
)

func main() {
	srv := server.New()
	srv.Setup()
	defer srv.Teardown()
	srv.Listen()
	if err := srv.Error(); err != nil {
		log.Fatal(err)
	}
}
