/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package middleware

import (
	"fmt"
	"time"

	"github.com/archnum/sdk.base/logger"
	"github.com/archnum/sdk.http/api/context"
	"github.com/archnum/sdk.http/api/core"
)

func Logger(logger *logger.Logger) func(core.Handler) core.Handler {
	return func(next core.Handler) core.Handler {
		return core.HandlerFunc(
			func(ctx context.Context) error {
				// TODO
				t0 := time.Now()
				defer func() {
					d := time.Since(t0)
					fmt.Println(d)
				}()

				return next.Serve(ctx)
			},
		)
	}
}

/*
####### END ############################################################################################################
*/
