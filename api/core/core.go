/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package core

import "github.com/archnum/sdk.http/api/render"

type (
	Handler interface {
		Serve(rr render.Renderer) error
	}

	HandlerFunc    func(rr render.Renderer) error
	MiddlewareFunc func(handler Handler) Handler
)

func (fn HandlerFunc) Serve(rr render.Renderer) error {
	return fn(rr)
}

/*
####### END ############################################################################################################
*/
