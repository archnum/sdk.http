/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package bind

import (
	"net/http"
	"strconv"

	"github.com/archnum/sdk.base/kv"
	"github.com/archnum/sdk.http/api/failure"
	"github.com/archnum/sdk.http/api/render"
)

func QueryInt(rr render.Renderer, name string, _defaultValue int) (int, error) {
	s := rr.Request().URL.Query().Get(name)
	if s == "" {
		return _defaultValue, nil
	}

	value, err := strconv.Atoi(s)
	if err != nil {
		return 0,
			failure.New( ///////////////////////////////////////////////////////////////////////////////////////////////
				http.StatusBadRequest,
				"bad type for this query parameter",
				kv.String("name", name),
			)
	}

	return value, nil
}

func QueryString(rr render.Renderer, name, _defaultValue string) string {
	value := rr.Request().URL.Query().Get(name)
	if value == "" {
		return _defaultValue
	}

	return value
}

func QueryStringSlice(rr render.Renderer, name string, _defaultValue ...string) []string {
	if value, ok := rr.Request().URL.Query()[name]; ok {
		return value
	}

	return _defaultValue
}

/*
####### END ############################################################################################################
*/
