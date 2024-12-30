/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/archnum/sdk.http/api/context"
	"github.com/archnum/sdk.http/api/core"
	"github.com/archnum/sdk.http/api/failure"
)

type (
	Manager interface {
		http.Handler
		Router() Router
	}

	implManager struct {
		router           *implRouter
		notFound         func() core.Handler
		methodNotAllowed func(allowedMethods []string) core.Handler
	}
)

func New(p *Params) *implManager {
	p.fix()

	return &implManager{
		router:           newRouter(newSegment()),
		notFound:         p.NotFound,
		methodNotAllowed: p.MethodNotAllowed,
	}
}

func (impl *implManager) Router() Router {
	return impl.router
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
			serve(wrap(mws, impl.notFound()), ctx) //------------------------------------------------------- 404 -------
			return
		}

		mws = append(mws, seg.middlewares...)
	}

	fn, ok := seg.fns[r.Method]
	if !ok {
		serve(wrap(mws, impl.methodNotAllowed(seg.allowedMethods())), ctx) //------------------------------- 405 -------
		return
	}

	serve(wrap(mws, fn), ctx)
}

func notFound() core.Handler {
	return core.HandlerFunc(
		func(ctx context.Context) error {
			ctx.ResponseWriter().WriteHeader(http.StatusNotFound)
			return nil
		},
	)
}

func methodNotAllowed(allowedMethods []string) core.Handler {
	return core.HandlerFunc(
		func(ctx context.Context) error {
			if len(allowedMethods) > 0 {
				ctx.ResponseWriter().Header().Set("Allow", strings.Join(allowedMethods, ", "))
			}

			ctx.ResponseWriter().WriteHeader(http.StatusMethodNotAllowed)

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

	f := new(failure.Failure)

	if !errors.As(err, f) {
		f = failure.New(http.StatusInternalServerError, err.Error())
	}

	ctx.WriteError(f)
}

/*
####### END ############################################################################################################
*/
