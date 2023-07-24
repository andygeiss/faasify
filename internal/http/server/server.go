package server

import (
	"crypto/tls"
	"net/http"
	"time"

	"golang.org/x/crypto/acme/autocert"
)

func ListenAndServe(domain, mode, url string) error {
	// save the domain/url
	Domain = domain
	Url = url
	// Create a custom http.Server with timeouts
	srv := &http.Server{
		Addr:         ":80",
		Handler:      http.TimeoutHandler(router(), time.Second*5, ""),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	// Run the server
	switch mode {
	case "prod":
		certManager := autocert.Manager{
			Cache:      autocert.DirCache("."),
			HostPolicy: autocert.HostWhitelist(domain),
			Prompt:     autocert.AcceptTOS,
		}
		srv.Addr = ":443"
		srv.TLSConfig = &tls.Config{
			GetCertificate: certManager.GetCertificate,
		}
		go http.ListenAndServe(":80", certManager.HTTPHandler(nil))
		return srv.ListenAndServeTLS("", "")
	default:
		return srv.ListenAndServe()
	}
}
