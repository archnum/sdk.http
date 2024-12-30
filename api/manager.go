/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package api

import (
	"net/http"
	"strings"

	"github.com/archnum/sdk.http/api/context"
	"github.com/archnum/sdk.http/api/core"
)

type (
	Manager interface {
		http.Handler
		Router() Router
	}

	implManager struct {
		router *implRouter
	}
)

func New() *implManager {
	return &implManager{
		router: newRouter(newSegment()),
	}
}

func (impl *implManager) Router() Router {
	return impl.router
}

func notFound(w http.ResponseWriter) core.Handler {
	return core.HandlerFunc(
		func(_ context.Context) error {
			w.WriteHeader(http.StatusNotFound)
			return nil
		},
	)
}

func methodNotAllowed(w http.ResponseWriter, seg *segment) core.Handler {
	return core.HandlerFunc(
		func(_ context.Context) error {
			allowedMethods := seg.allowedMethods()

			if len(allowedMethods) > 0 {
				w.Header().Set("Allow", strings.Join(allowedMethods, ", "))
			}

			w.WriteHeader(http.StatusMethodNotAllowed)

			return nil
		},
	)
}

func wrap(middlewares []core.MiddlewareFunc, handler core.Handler) core.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	return handler
}

func serve(handler core.Handler, ctx context.Context) {
	err := handler.Serve(ctx)
	if err == nil {
		return
	}

	// TODO
}

func (impl *implManager) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var ok bool

	ctx := context.New(w, r)
	seg := impl.router.seg

	mws := make([]core.MiddlewareFunc, 0, 10)
	mws = append(mws, seg.middlewares...)

	for _, s := range strings.Split(r.URL.EscapedPath(), "/") {
		if s == "" {
			continue
		}

		seg, ok = seg.nextSegment(ctx, s)
		if !ok {
			serve(wrap(mws, notFound(w)), ctx) //----------------------------------------------------------- 404 -------
			return
		}

		mws = append(mws, seg.middlewares...)
	}

	fn, ok := seg.fns[r.Method]
	if !ok {
		serve(wrap(mws, methodNotAllowed(w, seg)), ctx) //-------------------------------------------------- 405 -------
		return
	}

	serve(wrap(mws, fn), ctx)
}

/*
####### END ############################################################################################################
*/
