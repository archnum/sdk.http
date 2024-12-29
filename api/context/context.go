/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package context

import "net/http"

type (
	Context interface {
		AddURLParam(name, value string)
		URLParam(name string) (string, bool)
	}

	implContext struct {
		response http.ResponseWriter
		request  *http.Request
		params   map[string]string
	}
)

func New(w http.ResponseWriter, r *http.Request) *implContext {
	return &implContext{
		response: w,
		request:  r,
		params:   make(map[string]string),
	}
}

func (ctx *implContext) AddURLParam(name, value string) {
	ctx.params[name] = value
}

func (ctx *implContext) URLParam(name string) (string, bool) {
	value, ok := ctx.params[name]
	return value, ok
}

/*
####### END ############################################################################################################
*/
