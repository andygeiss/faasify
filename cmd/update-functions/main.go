package main

import (
	"log"

	"github.com/andygeiss/faasify/internal/http/generator"
)

func main() {
	gen := generator.New()
	gen.Setup()
	if err := gen.Error(); err != nil {
		log.Fatal(err)
	}
}
