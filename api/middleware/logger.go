/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package middleware

import (
	"fmt"
	"time"

	"github.com/archnum/sdk.base/logger"
	"github.com/archnum/sdk.http/api/core"
	"github.com/archnum/sdk.http/api/render"
)

func Logger(logger *logger.Logger) func(core.Handler) core.Handler {
	return func(next core.Handler) core.Handler {
		return core.HandlerFunc(
			func(rr render.Renderer) error {
				// TODO
				t0 := time.Now()
				defer func() {
					d := time.Since(t0)
					fmt.Println(d)
				}()

				return next.Serve(rr)
			},
		)
	}
}

/*
####### END ############################################################################################################
*/
