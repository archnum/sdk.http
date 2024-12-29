/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package api

import (
	"net/http"
	"strings"
)

type (
	Router interface {
		Mount(pattern string, fn func(router Router))
		Get(pattern string, fn HandlerFunc)
		Post(pattern string, fn HandlerFunc)
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

func (impl *implRouter) Mount(pattern string, fn func(router Router)) {
	seg := impl.seg

	for _, s := range strings.Split(pattern, "/") {
		if s == "" {
			continue
		}

		tmp, ok := seg.childs[s]
		if ok {
			seg = tmp
		} else {
			seg.childs[s] = newSegment()
			seg = seg.childs[s]
		}
	}

	fn(newRouter(seg))
}

func (impl *implRouter) handle(method string, pattern string, fn HandlerFunc) {
	seg := impl.seg

	for _, s := range strings.Split(pattern, "/") {
		if s == "" {
			continue
		}

		tmp, ok := seg.childs[s]
		if ok {
			seg = tmp
		} else {
			seg.childs[s] = newSegment()
			seg = seg.childs[s]
		}
	}

	seg.addHandlerFunc(method, fn)
}

func (impl *implRouter) Get(pattern string, fn HandlerFunc) {
	impl.handle(http.MethodGet, pattern, fn)
}

func (impl *implRouter) Post(pattern string, fn HandlerFunc) {
	impl.handle(http.MethodPost, pattern, fn)
}

/*
####### END ############################################################################################################
*/
