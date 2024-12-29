/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package api

import (
	"github.com/archnum/sdk.http/api/context"
	"github.com/archnum/sdk.http/api/core"
)

type (
	segment struct {
		childs      map[string]*segment
		fns         map[string]core.HandlerFunc
		param       string
		middlewares []core.MiddlewareFunc
	}
)

func newSegment() *segment {
	return &segment{
		childs: make(map[string]*segment),
	}
}

func (seg *segment) addMiddlewares(middlewares ...core.MiddlewareFunc) {
	seg.middlewares = append(seg.middlewares, middlewares...)
}

func (seg *segment) addHandlerFunc(method string, fn core.HandlerFunc) {
	if seg.fns == nil {
		seg.fns = make(map[string]core.HandlerFunc)
	}

	seg.fns[method] = fn
}

func (seg *segment) nextSegment(ctx context.Context, s string) (*segment, bool) {
	tmp, ok := seg.childs[s]
	if ok {
		return tmp, true
	}

	if seg.param != "" {
		ctx.AddURLParam(seg.param, s)
		return seg, true
	}

	return nil, false
}

/*
####### END ############################################################################################################
*/
