/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package server

import (
	"crypto/tls"
	"crypto/x509"
	"os"
	"time"

	"github.com/archnum/sdk.base/failure"
	"github.com/archnum/sdk.base/kv"
)

const (
	_defaultAddr = ":5635"
)

type (
	Config struct {
		Addr         string        `dm:"addr"`
		CertFile     string        `dm:"cert_file"`
		KeyFile      string        `dm:"key_file"`
		CAFile       string        `dm:"ca_file"`
		IdleTimeout  time.Duration `dm:"idle_timeout"`
		ReadTimeout  time.Duration `dm:"read_timeout"`
		WriteTimeout time.Duration `dm:"write_timeout"`
	}
)

func (cfg *Config) validate() error {
	if cfg.Addr == "" {
		cfg.Addr = _defaultAddr
	}

	return nil
}

func (cfg *Config) setupTLS() (*tls.Config, error) {
	if cfg.CertFile == "" && cfg.KeyFile == "" {
		return nil, nil
	}

	cert, err := tls.LoadX509KeyPair(cfg.CertFile, cfg.KeyFile)
	if err != nil {
		return nil,
			failure.WithMessage( ///////////////////////////////////////////////////////////////////////////////////////
				err,
				"failed to load certificate or key file",
				kv.String("cert_file", cfg.CertFile),
				kv.String("key_file", cfg.KeyFile),
			)
	}

	tlsConfig := &tls.Config{
		MinVersion:               tls.VersionTLS13,
		PreferServerCipherSuites: true,
		Certificates:             []tls.Certificate{cert},
	}

	if cfg.CAFile != "" {
		bs, err := os.ReadFile(cfg.CAFile)
		if err != nil {
			return nil, err
		}

		ca := x509.NewCertPool()
		ok := ca.AppendCertsFromPEM(bs)
		if !ok {
			return nil,
				failure.New("failed to load certificate authority", kv.String("file", cfg.CAFile)) /////////////////////
		}

		tlsConfig.ClientCAs = ca
		tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
	}

	return tlsConfig, nil
}

/*
####### END ############################################################################################################
*/
