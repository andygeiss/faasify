package server

import (
	"log"
	"net/http"
	"os"
	"time"
)

var securityToken string

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

	addr := os.Getenv("FAASIFY_ADDRESS")
	if addr == "" {
		addr = ":3000"
	}
	a.addr = addr

	a.srv = &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 5,
		Handler:      http.TimeoutHandler(router(), time.Second*1, ""),
	}

	securityToken = os.Getenv("FAASIFY_TOKEN")
	if securityToken == "" {
		securityToken = "TOKEN_SHOULD_BE_CHANGED"
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
