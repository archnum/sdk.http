/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package api

import "github.com/archnum/sdk.http/api/core"

type (
	Params struct {
		NotFound         core.Handler
		MethodNotAllowed core.Handler
	}
)

func (p *Params) fix() {
	if p.NotFound == nil {
		p.NotFound = notFound()
	}

	// TODO
}

/*
####### END ############################################################################################################
*/
