// Copyright 2017, Project ArteMisc
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package gitter

import (
	"fmt"
	"net/http"
)

// Server
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
	return fmt.Sprintf("%s:%d", s.c.Host, s.c.Port)
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
func (s *Server) HttpServer() (srv *http.Server) {
	srv = &http.Server{
		Addr:    s.Address(),
		Handler: s.Handler(),
	}
	if s.c.Tls != nil {
		srv.TLSConfig = s.c.Tls.Config
	}
	return
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
