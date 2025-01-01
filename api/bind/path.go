/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package bind

import (
	"net/http"
	"strconv"

	"github.com/archnum/sdk.base/kv"
	"github.com/archnum/sdk.base/uuid"
	"github.com/archnum/sdk.http/api/failure"
	"github.com/archnum/sdk.http/api/render"
)

func errBadType(name, expected, value string) error {
	return failure.New( ////////////////////////////////////////////////////////////////////////////////////////////////
		http.StatusBadRequest,
		"bad type for this path parameter",
		kv.String("name", name),
		kv.String("expected", expected),
		kv.String("value", value),
	)
}

func PathInt(rr render.Renderer, name string) (int, error) {
	s, err := PathString(rr, name)
	if err != nil {
		return 0, err
	}

	if value, err := strconv.Atoi(s); err == nil {
		return value, nil
	}

	return 0, errBadType(name, "int", s)
}

func PathString(rr render.Renderer, name string) (string, error) {
	value, ok := rr.URLParam(name)
	if !ok {
		return "",
			failure.New( ///////////////////////////////////////////////////////////////////////////////////////////////
				http.StatusBadRequest,
				"this path parameter doesn't exist",
				kv.String("name", name),
			)
	}

	return value, nil
}

func PathUUID(rr render.Renderer, name string) (uuid.UUID, error) {
	s, err := PathString(rr, name)
	if err != nil {
		return "", err
	}

	if value, ok := uuid.ConvertString(s); ok {
		return value, nil
	}

	return "", errBadType(name, "UUID", s)
}

/*
####### END ############################################################################################################
*/
