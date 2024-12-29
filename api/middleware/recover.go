/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package middleware

import (
	"fmt"

	"github.com/archnum/sdk.base/logger"
	"github.com/archnum/sdk.http/api/context"
	"github.com/archnum/sdk.http/api/core"
)

func Recover(logger *logger.Logger) func(core.Handler) core.Handler {
	return func(next core.Handler) core.Handler {
		return core.HandlerFunc(
			func(ctx context.Context) error {
				fmt.Println("RECOVER middleware")
				return next.Serve(ctx)
			},
		)
	}
}

/*
####### END ############################################################################################################
*/
