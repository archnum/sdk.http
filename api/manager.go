/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package api

import (
	"net/http"
	"strings"

	"github.com/archnum/sdk.base/logger"
	"github.com/archnum/sdk.http/api/core"
	"github.com/archnum/sdk.http/api/render"
)

type (
	Manager interface {
		http.Handler
		Logger() *logger.Logger
		Router() Router
	}

	implManager struct {
		logger           *logger.Logger
		router           *implRouter
		notFound         func() core.Handler
		methodNotAllowed func(allowedMethods []string) core.Handler
	}
)

func New(p *Params) *implManager {
	p.fix()

	return &implManager{
		logger:           p.Logger,
		router:           newRouter(newSegment()),
		notFound:         p.NotFound,
		methodNotAllowed: p.MethodNotAllowed,
	}
}

func (impl *implManager) Logger() *logger.Logger {
	return impl.logger
}

func (impl *implManager) Router() Router {
	return impl.router
}

func (impl *implManager) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var ok bool

	rr := render.New(impl.logger, w, r)
	seg := impl.router.seg

	mws := make([]core.MiddlewareFunc, 0, 10)
	mws = append(mws, seg.middlewares...)

	s, path := splitPath(r.URL.EscapedPath())

	for s != "" {
		seg, ok = seg.nextSegment(rr, s)
		if !ok {
			serve(wrap(mws, impl.notFound()), rr) ////////////////////////////////////////////////////////// 404 ///////
			return
		}

		mws = append(mws, seg.middlewares...)

		s, path = splitPath(path)
	}

	fn, ok := seg.fns[r.Method]
	if !ok {
		serve(wrap(mws, impl.methodNotAllowed(seg.allowedMethods())), rr) ////////////////////////////////// 405 ///////
		return
	}

	serve(wrap(mws, fn), rr)
}

func notFound() core.Handler {
	return core.HandlerFunc(
		func(rr render.Renderer) error {
			rr.ResponseWriter().WriteHeader(http.StatusNotFound)
			return nil
		},
	)
}

func methodNotAllowed(allowedMethods []string) core.Handler {
	return core.HandlerFunc(
		func(rr render.Renderer) error {
			if len(allowedMethods) > 0 {
				rr.ResponseWriter().Header().Set("Allow", strings.Join(allowedMethods, ", "))
			}

			rr.ResponseWriter().WriteHeader(http.StatusMethodNotAllowed)

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

func serve(handler core.Handler, rr render.Renderer) {
	err := handler.Serve(rr)
	if err == nil {
		return
	}

	rr.WriteError(err)
}

/*
####### END ############################################################################################################
*/
