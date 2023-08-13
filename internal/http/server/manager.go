package server

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/andygeiss/faasify/internal/account"
	"github.com/andygeiss/faasify/internal/config"
	"golang.org/x/crypto/acme/autocert"
)

type Manager struct {
	err error
	cfg *config.Config
}

func (a *Manager) Error() error {
	return a.err
}

func (a *Manager) ListenAndServe() {
	if a.err != nil {
		return
	}
	// Create a custom http.Server with timeouts
	srv := &http.Server{
		Addr:         ":3000",
		Handler:      http.TimeoutHandler(router(a.cfg), 10*time.Second, ""),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	// Run the server
	var err error
	switch a.cfg.Mode {
	case "production":
		certManager := autocert.Manager{
			Cache:      autocert.DirCache("."),
			HostPolicy: autocert.HostWhitelist(a.cfg.Domain),
			Prompt:     autocert.AcceptTOS,
		}
		srv.Addr = ":443"
		srv.TLSConfig = &tls.Config{
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			},
			CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			GetCertificate:           certManager.GetCertificate,
			MinVersion:               tls.VersionTLS12,
			PreferServerCipherSuites: true,
		}
		go http.ListenAndServe(":3000", certManager.HTTPHandler(nil))
		err = srv.ListenAndServeTLS("", "")
	default:
		// create a default user in development mode
		a.cfg.AccountAccess.CreateAccount("faasify", a.cfg.Token)
		err = srv.ListenAndServeTLS("data/localhost.crt", "data/localhost.key")
	}
	if err != nil {
		a.err = err
	}
}

func (a *Manager) WithAccountAccess(accountAccess account.Access) *Manager {
	a.cfg.AccountAccess = accountAccess
	return a
}

func (a *Manager) WithAppName(appName string) *Manager {
	a.cfg.AppName = appName
	return a
}

func (a *Manager) WithDomain(domain string) *Manager {
	a.cfg.Domain = domain
	return a
}

func (a *Manager) WithMode(mode string) *Manager {
	a.cfg.Mode = mode
	return a
}

func (a *Manager) WithUrl(url string) *Manager {
	a.cfg.Url = url
	return a
}

func NewManager() *Manager {
	return &Manager{
		cfg: &config.Config{},
	}
}
