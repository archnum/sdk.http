/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package bind

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/archnum/sdk.base/kv"
	"github.com/archnum/sdk.http/api/apierr"
)

func decodeJSON(body io.ReadCloser, maxSize int64, to any) error {
	decoder := json.NewDecoder(body)

	if err := decoder.Decode(to); err != nil {
		var (
			syntaxError        *json.SyntaxError
			unmarshalTypeError *json.UnmarshalTypeError
		)

		switch {
		case errors.As(err, &syntaxError):
			return apierr.New(http.StatusBadRequest, syntaxError.Error()) //////////////////////////////////////////////

		case errors.Is(err, io.ErrUnexpectedEOF):
			return apierr.New(http.StatusBadRequest, "request body contains badly-formed JSON") ////////////////////////

		case errors.As(err, &unmarshalTypeError):
			return apierr.New( /////////////////////////////////////////////////////////////////////////////////////////
				http.StatusBadRequest,
				"request body contains an invalid value for a field",
				kv.String("value", unmarshalTypeError.Value),
				kv.String("field", unmarshalTypeError.Field),
			)

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			return apierr.New( /////////////////////////////////////////////////////////////////////////////////////////
				http.StatusBadRequest,
				"request body contains an unknown field",
				kv.String("field", strings.TrimPrefix(err.Error(), "json: unknown field ")),
			)

		case errors.Is(err, io.EOF):
			return apierr.New(http.StatusBadRequest, "request body must not be empty") /////////////////////////////////

		case err.Error() == "http: request body too large":
			return apierr.New( /////////////////////////////////////////////////////////////////////////////////////////
				http.StatusRequestEntityTooLarge,
				"the request body is too large",
				kv.Int64("max_size", maxSize),
			)

		default:
			return apierr.New(http.StatusInternalServerError, err.Error()) /////////////////////////////////////////////
		}
	}

	if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return apierr.New( /////////////////////////////////////////////////////////////////////////////////////////////
			http.StatusBadRequest,
			"request body must only contain a single JSON object",
		)
	}

	return nil
}

/*
####### END ############################################################################################################
*/
