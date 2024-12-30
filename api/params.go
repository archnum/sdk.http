/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package api

import "github.com/archnum/sdk.http/api/core"

type (
	Params struct {
		NotFound         func() core.Handler
		MethodNotAllowed func(allowedMethods []string) core.Handler
	}
)

func (p *Params) fix() {
	if p.NotFound == nil {
		p.NotFound = notFound
	}

	if p.MethodNotAllowed == nil {
		p.MethodNotAllowed = methodNotAllowed
	}
}

/*
####### END ############################################################################################################
*/
