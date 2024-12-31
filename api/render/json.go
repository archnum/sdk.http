/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package render

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (impl *implRenderer) writeJSON(status int, data any) error {
	buf := new(bytes.Buffer)

	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)

	if err := encoder.Encode(data); err != nil {
		http.Error(impl.ResponseWriter(), err.Error(), http.StatusInternalServerError) /////////////////////////////////
		return err
	}

	impl.setContentType()
	impl.ResponseWriter().WriteHeader(status)

	_, err := impl.ResponseWriter().Write(buf.Bytes())
	return err
}

/*
####### END ############################################################################################################
*/
