/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package api

import (
	"strings"

	"github.com/archnum/sdk.http/api/context"
	"github.com/archnum/sdk.http/api/core"
)

const (
	_paramPrefix = ":"
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

func (seg *segment) maybeSetParam(s string) {
	if param := strings.TrimPrefix(s, _paramPrefix); param != s {
		seg.param = param
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
		return seg.childs[_paramPrefix+seg.param], true
	}

	return nil, false
}

func (seg *segment) allowedMethods() []string {
	list := make([]string, 0, len(seg.fns))

	for method := range seg.fns {
		list = append(list, method)
	}

	return list
}

/*
####### END ############################################################################################################
*/
