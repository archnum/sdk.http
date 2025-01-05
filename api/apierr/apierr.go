/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package apierr

import (
	"errors"
	"net/http"

	"github.com/archnum/sdk.base/failure"
	"github.com/archnum/sdk.base/kv"
)

type (
	WithStatus struct {
		error
		status int
	}
)

func New(status int, msg string, kvs ...kv.KeyValue) *WithStatus {
	return &WithStatus{
		error:  failure.New(msg, kvs...),
		status: status,
	}
}

func WithMessage(status int, cause error, msg string, kvs ...kv.KeyValue) *WithStatus {
	return &WithStatus{
		error:  failure.WithMessage(cause, msg, kvs...),
		status: status,
	}
}

func WithError(status int, err error) *WithStatus {
	if err == nil {
		return nil
	}

	var ae *WithStatus
	if errors.As(err, &ae) {
		return ae
	}

	return &WithStatus{
		error:  err,
		status: status,
	}
}

func BadRequest(err error) *WithStatus {
	return WithError(http.StatusBadRequest, err)
}

func NotFound(err error) *WithStatus {
	return WithError(http.StatusNotFound, err)
}

func InternalServerError(err error) *WithStatus {
	return WithError(http.StatusInternalServerError, err)
}

func (ws *WithStatus) Unwrap() error {
	return ws.error
}

func (ws *WithStatus) Status() int {
	return ws.status
}

/*
####### END ############################################################################################################
*/
