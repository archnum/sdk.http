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

func wrap(middlewares []core.MiddlewareFunc, handler core.Handler) core.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	return handler
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
			// TODO
			w.WriteHeader(404) // AFAC
			return
		}

		mws = append(mws, seg.middlewares...)
	}

	fn, ok := seg.fns[r.Method]
	if !ok {
		// TODO
		w.WriteHeader(405) // AFAC
		return
	}

	err := wrap(mws, fn).Serve(ctx)
	if err == nil {
		w.WriteHeader(204) // AFAC
		return
	}

	// TODO

	w.WriteHeader(500) // AFAC
}

/*
####### END ############################################################################################################
*/
