/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package api

import (
	"net/http"

	"github.com/archnum/sdk.http/api/core"
)

type (
	Router interface {
		Use(middlewares ...core.MiddlewareFunc)
		Mount(pattern string, fn func(router Router))
		Delete(pattern string, fn core.HandlerFunc)
		Get(pattern string, fn core.HandlerFunc)
		Options(pattern string, fn core.HandlerFunc)
		Patch(pattern string, fn core.HandlerFunc)
		Post(pattern string, fn core.HandlerFunc)
		Put(pattern string, fn core.HandlerFunc)
	}

	implRouter struct {
		seg *segment
	}
)

func newRouter(seg *segment) *implRouter {
	return &implRouter{
		seg: seg,
	}
}

func (impl *implRouter) Use(middlewares ...core.MiddlewareFunc) {
	impl.seg.addMiddlewares(middlewares...)
}

func (impl *implRouter) Mount(pattern string, fn func(router Router)) {
	seg := impl.seg.buildTree(pattern)
	fn(newRouter(seg))
}

func (impl *implRouter) handle(method string, pattern string, fn core.HandlerFunc) {
	seg := impl.seg.buildTree(pattern)
	seg.addHandlerFunc(method, fn)
}

func (impl *implRouter) Delete(pattern string, fn core.HandlerFunc) {
	impl.handle(http.MethodDelete, pattern, fn)
}

func (impl *implRouter) Get(pattern string, fn core.HandlerFunc) {
	impl.handle(http.MethodGet, pattern, fn)
}

func (impl *implRouter) Options(pattern string, fn core.HandlerFunc) {
	impl.handle(http.MethodOptions, pattern, fn)
}

func (impl *implRouter) Patch(pattern string, fn core.HandlerFunc) {
	impl.handle(http.MethodPatch, pattern, fn)
}

func (impl *implRouter) Post(pattern string, fn core.HandlerFunc) {
	impl.handle(http.MethodPost, pattern, fn)
}

func (impl *implRouter) Put(pattern string, fn core.HandlerFunc) {
	impl.handle(http.MethodPut, pattern, fn)
}

/*
####### END ############################################################################################################
*/
