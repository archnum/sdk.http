/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package util

import "net/http"

const (
	HeaderXRequestID = "X-Request-Id"
)

func RequestID(r *http.Request) string {
	return r.Header.Get(HeaderXRequestID)
}

func SetRequestID(r *http.Request, id string) {
	r.Header.Set(HeaderXRequestID, id)
}

/*
####### END ############################################################################################################
*/
