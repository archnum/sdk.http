/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package server

import (
	"log"
	"net/http"

	"github.com/archnum/sdk.base/failure"
)

type (
	Params struct {
		Config  *Config
		Handler http.Handler // Required
		Logger  *log.Logger
	}
)

func (p *Params) validate() error {
	if p.Config == nil {
		p.Config = new(Config)
	} else if err := p.Config.validate(); err != nil {
		return err
	}

	if p.Handler == nil {
		return failure.New("parameter 'Handler' is required") //////////////////////////////////////////////////////////
	}

	return nil
}

/*
####### END ############################################################################################################
*/
