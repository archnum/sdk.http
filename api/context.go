/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package api

import "net/http"

type (
	Context interface {
	}

	implContext struct {
		response http.ResponseWriter
		request  *http.Request
	}
)

func newContext(w http.ResponseWriter, r *http.Request) *implContext {
	return &implContext{
		response: w,
		request:  r,
	}
}

/*
####### END ############################################################################################################
*/
