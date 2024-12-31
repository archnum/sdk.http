/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package render

import (
	"net/http"

	"github.com/archnum/sdk.base/kv"
	"github.com/archnum/sdk.base/logger"
	"github.com/archnum/sdk.http/api/util"
)

type (
	Renderer interface {
		http.ResponseWriter
		Request() *http.Request
		AddURLParam(name, value string)
		URLParam(name string) (string, bool)
		WriteData(status int, data any)
		WriteError(err ErrorWithStatus)
	}

	// TODO: un pool ?
	// fieldalignment
	implRenderer struct {
		http.ResponseWriter
		request     *http.Request
		logger      *logger.Logger
		params      map[string]string
		contentType string
	}
)

func New(logger *logger.Logger, w http.ResponseWriter, r *http.Request) *implRenderer {
	return &implRenderer{
		ResponseWriter: w,
		request:        r,
		logger:         logger,
		params:         make(map[string]string),
		contentType:    util.ContentTypeJSON,
	}
}

func (impl *implRenderer) Request() *http.Request {
	return impl.request
}

func (impl *implRenderer) AddURLParam(name, value string) {
	impl.params[name] = value
}

func (impl *implRenderer) URLParam(name string) (string, bool) {
	value, ok := impl.params[name]
	return value, ok
}

func (impl *implRenderer) setContentType() {
	impl.Header().Set("Content-Type", impl.contentType)
}

func (impl *implRenderer) WriteData(status int, data any) {
	var err error

	switch impl.contentType {
	case util.ContentTypeYAML:
		err = impl.writeYAML(status, data)
	case util.ContentTypeXML:
		err = impl.writeXML(status, data)
	default:
		err = impl.writeJSON(status, data)
	}

	if err != nil {
		if impl.logger != nil {
			impl.logger.Error( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
				"Failed to write HTTP(S) response",
				kv.String("request", util.RequestID(impl.request)),
				kv.Int("status", status),
				kv.Error(err),
			)
		}
	}
}

func (impl *implRenderer) WriteError(err ErrorWithStatus) {
	requestID := util.RequestID(impl.request)
	status := err.Status()

	if impl.logger != nil {
		impl.logger.Error( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
			"HTTP(S) request failed",
			kv.String("id", requestID),
			kv.Int("status", status),
			kv.Error(err),
		)
	}

	impl.WriteData(status, &dataError{Error: err.Error(), RequestID: requestID})
}

/*
####### END ############################################################################################################
*/
