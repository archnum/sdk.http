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

func (impl *implManager) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(w, r)
	seg := impl.router.seg
	var ok bool

	for _, s := range strings.Split(r.URL.EscapedPath(), "/") {
		if s == "" {
			continue
		}

		seg, ok = seg.childs[s]
		if !ok {
			break
		}
	}

	if seg != nil {
		if fn, ok := seg.fns[r.Method]; ok {
			_ = fn(ctx)
		}
	}

	w.WriteHeader(204) // TODO
}

/*
####### END ############################################################################################################
*/
