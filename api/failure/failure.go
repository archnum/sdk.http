/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package failure

import (
	"github.com/archnum/sdk.base/failure"
	"github.com/archnum/sdk.base/kv"
)

type (
	Failure struct {
		*failure.Failure
		status int
	}
)

func New(status int, msg string, kvs ...kv.KeyValue) *Failure {
	return &Failure{
		Failure: failure.New(msg, kvs...),
		status:  status,
	}
}

func WithMessage(status int, cause error, msg string, kvs ...kv.KeyValue) *Failure {
	return &Failure{
		Failure: failure.WithMessage(cause, msg, kvs...),
		status:  status,
	}
}

func (f *Failure) Status() int {
	return f.status
}

/*
####### END ############################################################################################################
*/
