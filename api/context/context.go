/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package context

import "net/http"

type (
	Context interface {
		ResponseWriter() http.ResponseWriter
		AddURLParam(name, value string)
		URLParam(name string) (string, bool)
	}

	implContext struct {
		responseWriter http.ResponseWriter
		request        *http.Request
		params         map[string]string
	}
)

func New(w http.ResponseWriter, r *http.Request) *implContext {
	return &implContext{
		responseWriter: w,
		request:        r,
		params:         make(map[string]string),
	}
}

func (ctx *implContext) ResponseWriter() http.ResponseWriter {
	return ctx.responseWriter
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
