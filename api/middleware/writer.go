/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package middleware

import (
	"net/http"
)

type (
	// TODO: un pool ?
	responseWriter struct {
		http.ResponseWriter
		writeable  bool
		written    int
		statusCode int
	}
)

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{
		ResponseWriter: w,
		writeable:      true,
	}
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	if rw.writeable {
		rw.statusCode = statusCode
		rw.ResponseWriter.WriteHeader(statusCode)
		rw.writeable = false
	}
}

func (rw *responseWriter) Write(buf []byte) (int, error) {
	rw.WriteHeader(http.StatusOK)

	n, err := rw.ResponseWriter.Write(buf)
	rw.written = n

	return n, err
}

func (rw *responseWriter) status() int {
	return rw.statusCode
}

/*
####### END ############################################################################################################
*/
