/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package server

import (
	"context"
	"net/http"
)

type (
	Server struct {
		server *http.Server
	}
)

func New(p *Params) (*Server, error) {
	if err := p.validate(); err != nil {
		return nil, err
	}

	cfg := p.Config

	tlsConfig, err := cfg.setupTLS()
	if err != nil {
		return nil, err
	}

	server := &Server{
		server: &http.Server{
			Addr:              cfg.Addr,
			IdleTimeout:       cfg.IdleTimeout,
			ReadTimeout:       cfg.ReadTimeout,
			WriteTimeout:      cfg.WriteTimeout,
			Handler:           p.Handler,
			ReadHeaderTimeout: 0,
			TLSConfig:         tlsConfig,
			ErrorLog:          p.Logger,
		},
	}

	return server, nil
}

func (s *Server) Addr() string {
	return s.server.Addr
}

func (s *Server) TLS() bool {
	return s.server.TLSConfig != nil
}

func (s *Server) Start() error {
	if s.TLS() {
		return s.server.ListenAndServeTLS("", "")
	}

	return s.server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	s.server.SetKeepAlivesEnabled(false)
	return s.server.Shutdown(ctx)
}

/*
####### END ############################################################################################################
*/
