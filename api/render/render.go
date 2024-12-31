/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package render

import (
	"errors"
	"net/http"

	"github.com/archnum/sdk.base/kv"
	"github.com/archnum/sdk.base/logger"

	"github.com/archnum/sdk.http/api/failure"
	"github.com/archnum/sdk.http/api/util"
)

type (
	Renderer interface {
		ResponseWriter() http.ResponseWriter
		SetResponseWriter(rw http.ResponseWriter)
		Request() *http.Request
		AddURLParam(name, value string)
		URLParam(name string) (string, bool)
		WriteData(status int, data any)
		WriteError(err error)
	}

	// TODO: un pool ?
	// fieldalignment
	implRenderer struct {
		responseWriter http.ResponseWriter
		request        *http.Request
		logger         *logger.Logger
		params         map[string]string
		contentType    string
	}
)

func New(logger *logger.Logger, w http.ResponseWriter, r *http.Request) *implRenderer {
	return &implRenderer{
		responseWriter: w,
		request:        r,
		logger:         logger,
		params:         make(map[string]string),
		contentType:    util.ContentTypeJSON,
	}
}

func (impl *implRenderer) ResponseWriter() http.ResponseWriter {
	return impl.responseWriter
}

func (impl *implRenderer) SetResponseWriter(rw http.ResponseWriter) {
	impl.responseWriter = rw
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
	impl.responseWriter.Header().Set("Content-Type", impl.contentType)
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

func (impl *implRenderer) WriteError(err error) {
	f := new(failure.Failure)

	if !errors.As(err, f) {
		f = failure.New(http.StatusInternalServerError, err.Error())
	}

	requestID := util.RequestID(impl.request)
	status := f.Status()

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
