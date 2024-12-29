/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package core

import "github.com/archnum/sdk.http/api/context"

type (
	Handler interface {
		Serve(ctx context.Context) error
	}

	HandlerFunc    func(ctx context.Context) error
	MiddlewareFunc func(handler Handler) Handler
)

func (fn HandlerFunc) Serve(ctx context.Context) error {
	return fn(ctx)
}

/*
####### END ############################################################################################################
*/
