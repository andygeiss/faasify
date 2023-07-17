package server

import (
	"log"
	"net/http"
	"os"
	"time"
)

type server struct {
	addr  string
	err   error
	srv   *http.Server
	token string
}

func (a *server) Error() error {
	return a.err
}

func (a *server) Listen() {
	if a.err != nil {
		return
	}
	log.Printf("Start listening at %s ...", a.addr)
	if err := a.srv.ListenAndServe(); err != nil {
		a.err = err
	}
}

func (a *server) Setup() {
	if a.err != nil {
		return
	}
	// Configure the server address via environment
	addr := os.Getenv("FAASIFY_ADDRESS")
	if addr == "" {
		addr = ":3000"
	}
	a.addr = addr
	// Create a custom http.Server with timeouts
	a.srv = &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 5,
		Handler:      http.TimeoutHandler(router(), time.Second*1, ""),
	}
}

func (a *server) Teardown() {
	if a.err != nil {
		return
	}
}

func New() *server {
	return &server{}
}
