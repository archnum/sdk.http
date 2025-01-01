/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package middleware

import (
	"net/http"

	"github.com/archnum/sdk.base/kv"
	"github.com/archnum/sdk.base/logger"
	_base "github.com/archnum/sdk.base/util"

	"github.com/archnum/sdk.http/api/core"
	"github.com/archnum/sdk.http/api/render"
	"github.com/archnum/sdk.http/api/util"
)

func Recover(logger *logger.Logger) func(core.Handler) core.Handler {
	return func(next core.Handler) core.Handler {
		return core.HandlerFunc(
			func(rr render.Renderer) error {
				defer func() {
					if data := recover(); data != nil {
						if data == http.ErrAbortHandler {
							panic(data)
						}

						logger.Error( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
							"Request error recovered",
							kv.String("id", util.RequestID(rr.Request())),
							kv.Any("data", data),
							kv.String("stack", _base.Stack(5)),
						)

						rr.ResponseWriter().WriteHeader(http.StatusInternalServerError)
					}
				}()

				return next.Serve(rr)
			},
		)
	}
}

/*
####### END ############################################################################################################
*/
