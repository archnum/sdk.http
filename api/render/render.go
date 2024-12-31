/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package render

import "net/http"

type (
	Renderer interface {
		ResponseWriter() http.ResponseWriter
		Request() *http.Request
		AddURLParam(name, value string)
		URLParam(name string) (string, bool)
		WriteError(err ErrorWithStatus)
	}

	implRenderer struct {
		responseWriter http.ResponseWriter
		request        *http.Request
		params         map[string]string
	}
)

func New(w http.ResponseWriter, r *http.Request) *implRenderer {
	return &implRenderer{
		responseWriter: w,
		request:        r,
		params:         make(map[string]string),
	}
}

func (ctx *implRenderer) ResponseWriter() http.ResponseWriter {
	return ctx.responseWriter
}

func (ctx *implRenderer) Request() *http.Request {
	return ctx.request
}

func (ctx *implRenderer) AddURLParam(name, value string) {
	ctx.params[name] = value
}

func (ctx *implRenderer) URLParam(name string) (string, bool) {
	value, ok := ctx.params[name]
	return value, ok
}

func (ctx *implRenderer) WriteError(err ErrorWithStatus) {
	// TODO
	ctx.responseWriter.WriteHeader(500)
}

/*
####### END ############################################################################################################
*/
