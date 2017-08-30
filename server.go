package gitter

import (
	"net/http"
	"time"
)

const serverTimeout = time.Second * 20

type Server struct {
	c *Config
}

// New creates a new Server instance using the provided config.
func New(conf *Config) *Server {
	return &Server{conf}
}

// Address returns the address that the Server should listen on, if started
// using Server.Run().
func (s *Server) Address() string {
	return s.c.Host
}

// Handler returns a http.Handler that handles requests based on the Server's
// package configuration.
//
// The http.Handler is implemented as a *http.ServeMux, using the full package
// names (host and path) as routes.
func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()

	for _, pkg := range s.c.Packages {
		pkg.Handle(mux)
	}

	return mux
}

// HttpServer returns an *http.Server that serves the gitter's requests.
func (s *Server) HttpServer() *http.Server {
	return &http.Server{
		Addr:      s.Address(),
		Handler:   s.Handler(),
		TLSConfig: s.c.Tls.Config,
	}
}

// ListenAndServe starts the server in http (non-encrypted) mode.
func (s *Server) ListenAndServe() (err error) {
	err = s.HttpServer().ListenAndServe()
	return
}

// ListenAndServeTLS starts the server in https (encrypted) mode.
func (s *Server) ListenAndServeTLS() (err error) {
	err = s.HttpServer().ListenAndServeTLS(s.c.Tls.CertPath, s.c.Tls.KeyPath)
	return
}
