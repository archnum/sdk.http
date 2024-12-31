/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package middleware

import (
	"time"

	"github.com/archnum/sdk.base/kv"
	"github.com/archnum/sdk.base/logger"
	"github.com/archnum/sdk.base/uuid"

	"github.com/archnum/sdk.http/api/core"
	"github.com/archnum/sdk.http/api/render"
	"github.com/archnum/sdk.http/api/util"
)

func Logger(logger *logger.Logger) func(core.Handler) core.Handler {
	return func(next core.Handler) core.Handler {
		return core.HandlerFunc(
			func(rr render.Renderer) error {
				r := rr.Request()

				id := util.RequestID(r)
				if id == "" {
					var err error

					if id, err = uuid.String(); err == nil {
						util.SetRequestID(r, id)
					}
				}

				if logger == nil {
					return next.Serve(rr)
				}

				logger.Debug( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
					"HTTP(S) request",
					kv.String("id", id),
					kv.String("from", r.RemoteAddr),
					kv.String("method", r.Method),
					kv.String("uri", r.URL.RequestURI()),
					kv.Int64("content_length", r.ContentLength),
				)

				rw := newResponseWriter(rr.ResponseWriter())
				rr.SetResponseWriter(rw)

				t0 := time.Now()
				err := next.Serve(rr)
				d := time.Since(t0)

				if err != nil {
					rr.WriteError(err)
				}

				logger.Debug( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
					"HTTP(S) response",
					kv.String("id", id),
					kv.Int("status", rw.status()),
					kv.Duration("duration", d),
					kv.Int("written", rw.written),
				)

				return nil
			},
		)
	}
}

/*
####### END ############################################################################################################
*/
