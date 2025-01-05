/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package bind

import (
	"net/http"
	"strings"

	"github.com/archnum/sdk.base/kv"
	"github.com/archnum/sdk.http/api/apierr"
	"github.com/archnum/sdk.http/api/render"
)

const (
	DefaultMaxBodySize = 32 * 1024 // 32 Ko
)

func Body(rr render.Renderer, maxSize int64, to any) error {
	r := rr.Request()
	r.Body = http.MaxBytesReader(rr.ResponseWriter(), r.Body, maxSize)

	contentType := r.Header.Get("Content-Type")

	if strings.HasPrefix(contentType, "application/json") {
		return decodeJSON(r.Body, maxSize, to)
	}

	return apierr.New( /////////////////////////////////////////////////////////////////////////////////////////////////
		http.StatusUnsupportedMediaType,
		"Unsupported content Type",
		kv.String("content-type", contentType),
	)
}

/*
####### END ############################################################################################################
*/
